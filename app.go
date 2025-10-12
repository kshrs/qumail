package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"slices"
	"log"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"errors"
	"strings"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-message/mail"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	// "github.com/emersion/go-imap/client"
)

// App struct
type App struct {
	ctx context.Context
	mail *Mail
}

var (
	oauthConfig *oauth2.Config
	oauthToken  *oauth2.Token
	authURL     string
)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	
	// Load the dot env files
	godotenv.Load()
	usermail := os.Getenv("EMAIL")
	password := os.Getenv("APP_PASSWORD")

	// Connect with the mail client
	m, err := NewMail(usermail, password)
	if err != nil {
		log.Fatal(err)
	}
	
	err = m.Connect("imap.gmail.com:993")
	if err != nil {
		log.Fatal(err)
	}
	a.mail = m
	runtime.EventsEmit(a.ctx, "backend:ready")

	// OAuth Stuff
	oauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback", // Must match what you configure in Google Cloud (optional for desktop apps)
		Scopes:       []string{gmail.MailGoogleComScope},
		Endpoint:     google.Endpoint,
	}

	// Generate the URL for the user to visit
	// The "offline" access type is crucial for getting a refresh token
	authURL = oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)


}

// SignIn starts the OAuth2 login flow.
func (a *App) SignIn() (string, error) {
	// Channel to wait for the callback to complete
	callbackComplete := make(chan bool)

	// Start a temporary local web server
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// Get the authorization code from the query parameters
		code := r.URL.Query().Get("code")
		if code == "" {
			fmt.Fprintf(w, "Error: Could not find authorization code.")
			callbackComplete <- false
			return
		}

		// Exchange the code for a token
		token, err := oauthConfig.Exchange(context.Background(), code)
		if err != nil {
			fmt.Fprintf(w, "Error: Could not exchange code for token: %v", err)
			callbackComplete <- false
			return
		}

		// Store the token securely
		oauthToken = token
		// TODO: Save the token to a secure file here!
		log.Println("Successfully received token.")
		fmt.Fprintf(w, "<h1>Authentication Successful!</h1><p>You can now close this browser tab.</p>")
		
		// Signal that the process is complete
		callbackComplete <- true
	})

	// Run the server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Open the user's browser to the Google consent page
	runtime.BrowserOpenURL(a.ctx, authURL)

	// Wait for the callback to signal completion
	success := <-callbackComplete
	server.Shutdown(context.Background()) // Gracefully shut down the server

	if !success {
		return "", errors.New("authentication failed")
	}

	// Now you have the token, you can connect to IMAP/SMTP
	// For IMAP/SMTP, you must use a special XOAUTH2 mechanism
	// The access token is oauthToken.AccessToken
	// You would now re-initialize your IMAP client with this token
	
	return "Successfully signed in!", nil
}

func (a *App) FetchEmails(count uint32, section string) ([]Email, error) {
	if a.mail == nil {
		return nil, errors.New("not connected to the mail server")
	}

	valid_section_names := []string{"INBOX",
							"[Gmail]/Sent Mail",
							"[Gmail]/Spam",
							"[Gmail]/Trash",
							"[Gmail]/Drafts",
							"[Gmail]/All Mail",
							"[Gmail]/Starred",
							"[Gmail]/Important",
						}

	if !slices.Contains(valid_section_names, section) {
		section = "INBOX"
	}
		
	messages, err := a.mail.GetMessages(section, count)
	if err != nil {
		log.Fatal(err)
	}


	var emails []Email
	for _, msg := range messages {

		isRead := false
		for _, flag := range msg.Flags {
			if flag == imap.SeenFlag {
				isRead = true
				break
			}
		}
		isStarred := false
		for _, flag := range msg.Flags {
			if flag == imap.FlaggedFlag {
				isStarred = true
				break
			}
		}


		if msg.Envelope == nil {
			continue
		}

		from := ""
		if len(msg.Envelope.From) > 0 {
			addr := msg.Envelope.From[0]
			from = fmt.Sprintf("%s <%s@%s>",addr.PersonalName, addr.MailboxName, addr.HostName)
		}

		var subject string
		if msg.Envelope.Subject == "" {
			subject = "(No Subject)"
		} else {
			subject = msg.Envelope.Subject
		}

		email := Email{
			SeqNum: msg.SeqNum,
			From: from,
			Subject: subject,
			Date: msg.Envelope.Date.Format("02 Jan 2006 15:04"),
			IsRead: isRead,
			IsStarred: isStarred,
		}
		emails = append(emails, email)
	}
	
	return emails, nil

}

func cleanEmailList(emails string) []string {
	if emails == "" {
		return []string{}
	}

	list := strings.Split(emails, ",")
	for i, email := range list {
		list[i] = strings.TrimSpace(email)
	}

	return list
}

func (a *App) SendEmail(to, cc, bcc string, subject string, body string, attachmentPaths []string) (string, error) {
	if a.mail == nil {
		return "",errors.New("Connect to a server first")
	}

	toList := cleanEmailList(to)
	ccList := cleanEmailList(cc)
	bccList := cleanEmailList(bcc)


	err := a.mail.SendEmail(toList, ccList, bccList, subject, body, attachmentPaths)
	if err != nil {
		return "", err
	}
	return "Sent Email Successfully", nil
}

func (a *App) ReadEmail(seqNum uint32) (FullEmail, error) {
	if a.mail == nil || a.mail.Client == nil{
		return FullEmail{}, errors.New("Connect to a server first")
	}

	seqset := new(imap.SeqSet)
	seqset.AddNum(seqNum)
	
	items := []imap.FetchItem{imap.FetchBodyStructure, imap.FetchEnvelope}
	section := &imap.BodySectionName{}
	items = append(items, section.FetchItem())

	messages := make(chan *imap.Message, 1)
	if err := a.mail.Client.Fetch(seqset, items, messages); err != nil {
		return FullEmail{}, err
	}

	msg := <- messages
	if msg == nil {
		return FullEmail{}, errors.New("email not found")
	}

	var fullEmail FullEmail
	r := msg.GetBody(section)
	if r == nil {
		return FullEmail{}, errors.New("could not get message body")
	}

	mr, err := mail.CreateReader(r)
	if err != nil {
		return FullEmail{}, err

	}

	header := mr.Header
	fullEmail.From = header.Get("From")
	fullEmail.To = header.Get("To")
	fullEmail.Cc = header.Get("Cc")
	fullEmail.Subject = header.Get("Subject")

	for {
		p, err := mr.NextPart() 
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error reading message parts")
			continue
		}

		switch h := p.Header.(type) {
			case *mail.InlineHeader:
				contentType, _, _ := h.ContentType()
				if contentType == "text/html" {
					bodyBytes, _ :=  ioutil.ReadAll(p.Body)
					fullEmail.Body = string(bodyBytes)
				}
		}
	}

	return fullEmail, nil

}

func (a *App) PickFiles() ([]string, error) {
	return runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select files to attach",
	})
}

func (a *App) ToggleRead(seqNum uint32, currentStateIsRead bool) (string, error) {
	if a.mail == nil {
		return "",errors.New("Connect to a server first")
	}

	seqset := new(imap.SeqSet)
	seqset.AddNum(seqNum)

	var op imap.FlagsOp
	op = imap.AddFlags
	if currentStateIsRead {
		op = imap.RemoveFlags
	}
	item := imap.FormatFlagsOp(op, true)

	flag := imap.SeenFlag

	err := a.mail.Client.Store(seqset, item, []interface{}{flag}, nil)
	if err != nil {
		return "", err
	}

	return "Read Status Updated", nil
}

func (a *App) ToggleStarred(seqNum uint32, currentStateIsStarred bool) (string, error) {
	if a.mail == nil || a.mail.Client == nil {
		return "", errors.New("Connect to a server first")
	}

	seqset := new(imap.SeqSet)
	seqset.AddNum(seqNum)

	var op imap.FlagsOp
	op = imap.AddFlags
	if currentStateIsStarred {
		op = imap.RemoveFlags
	}
	item := imap.FormatFlagsOp(op, true)

	flag := imap.FlaggedFlag

	err := a.mail.Client.Store(seqset, item, []interface{}{flag}, nil)
	if err != nil {
		return "", err
	}

	return "Starred Status Updated", nil
}

func (a *App) shutdown(ctx context.Context) {
	if a.mail != nil {
		a.mail.Disconnect()
	}
}
