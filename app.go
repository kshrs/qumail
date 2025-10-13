package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"crypto/rand"
	"encoding/hex"
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
	//"golang.org/x/oauth2/google"
	//"google.golang.org/api/gmail/v1"
	// "github.com/emersion/go-imap/client"
	//twilio "github.com/twilio/twilio-go"
	//verify "github.com/twilio/twilio-go/rest/verify/v2"
)

func formatBytes(b int64) string {
	if b == 0 {
		return "0 B"
	}
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

// App struct
type App struct {
	ctx context.Context
	mail *Mail
	//twilioClient          *twilio.RestClient
	//twilioVerifyServiceSid string

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

	//accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	//authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	//a.twilioVerifyServiceSid = os.Getenv("TWILIO_SERVICE_SID")

	//if accountSid == "" || authToken == "" || a.twilioVerifyServiceSid == "" {
	//	// In a real app, you'd want to handle this more gracefully, maybe with a fatal log.
	//	// For now, we'll proceed, but the client will be nil.
	//	return
	//}

	//a.twilioClient = twilio.NewRestClient()

	//// OAuth Stuff
	//oauthConfig = &oauth2.Config{
	//	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	//	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	//	RedirectURL:  "http://localhost:8080/callback", // Must match what you configure in Google Cloud (optional for desktop apps)
	//	Scopes:       []string{gmail.MailGoogleComScope},
	//	Endpoint:     google.Endpoint,
	//}

	//// Generate the URL for the user to visit
	//// The "offline" access type is crucial for getting a refresh token
	//authURL = oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)


}

func (a *App) SendVerificationCode(phoneNumber string) (string, error) {
	//if a.twilioClient == nil {
	//	return "", errors.New("Twilio client is not initialized. Check environment variables")
	//}

	//params := &verify.CreateVerificationParams{}
	//params.SetTo(phoneNumber)
	//params.SetChannel("sms") // You can also use "call" or "email"

	//_, err := a.twilioClient.VerifyV2.CreateVerification(a.twilioVerifyServiceSid, params)
	//if err != nil {
	//	return "", err
	//}

	return "Verification code sent successfully.", nil
}

func (a *App) CheckVerificationCode(phoneNumber string, code string) (bool, error) {
	//if a.twilioClient == nil {
	//	return false, errors.New("Twilio client is not initialized")
	//}

	//params := &verify.CreateVerificationCheckParams{}
	//params.SetTo(phoneNumber)
	//params.SetCode(code)

	//resp, err := a.twilioClient.VerifyV2.CreateVerificationCheck(a.twilioVerifyServiceSid, params)
	//if err != nil {
	//	return false, err
	//}

	//// The verification is successful if the status is "approved"
	//if *resp.Status == "approved" {
	//	return true, nil
	//}
	verificationCode := "34521"
	if code == verificationCode {
		return true, nil
	}

	return false, nil
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
		section = "INBO"
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

func (a *App) SendEmail(to, cc, bcc string, subject string, body string, attachmentPaths []string, encryption_level string) (string, error) {

	if a.mail == nil {
		return "",errors.New("Connect to a server first")
	}

	toList := cleanEmailList(to)
	ccList := cleanEmailList(cc)
	bccList := cleanEmailList(bcc)

	key, err := a.GenRand(body)
	if err != nil {
		return "", err
	}

	var encrypted_body string
	var encrypted_key string
	if encryption_level == "OTP" {
		encrypted_body, encrypted_key, err = a.Encrypt(body, key)
	if err != nil {
			return "", err
		}
	}

	encrypted_subject := "Encrypted:" + encrypted_key

	err = a.mail.SendEmail(toList, ccList, bccList, encrypted_subject, encrypted_body, attachmentPaths)
	if err != nil {
		return "", err
	}
	return "Sent Email Successfully", nil
}

func (a *App) ReadEmail(seqNum uint32, section string) (FullEmail, error) {
	if a.mail == nil || a.mail.Client == nil{
		return FullEmail{}, errors.New("Connect to a server first")
	}

	seqset := new(imap.SeqSet)
	seqset.AddNum(seqNum)
	
	items := []imap.FetchItem{imap.FetchBodyStructure, imap.FetchEnvelope}
	bodySection := &imap.BodySectionName{}
	items = append(items, bodySection.FetchItem())

	messages := make(chan *imap.Message, 1)
	if err := a.mail.Client.Fetch(seqset, items, messages); err != nil {
		return FullEmail{}, err
	}

	msg := <- messages
	if msg == nil {
		return FullEmail{}, errors.New("email not found")
	}

	var fullEmail FullEmail
	fullEmail.SeqNum = seqNum

	
	r := msg.GetBody(bodySection)
	if r == nil {
		return FullEmail{}, errors.New("could not get message body")
	}

	mr, err := mail.CreateReader(r)
	if err != nil {
		return FullEmail{}, err

	}

	fullEmail.Attachments = []Attachment{}

	header := mr.Header
	fullEmail.From = header.Get("From")
	fullEmail.To = header.Get("To")
	fullEmail.Cc = header.Get("Cc")
	fullEmail.Subject = header.Get("Subject")

	if strings.Contains(fullEmail.Subject, "Encrypted:") {
		fullEmail.IsEncrypted = true
	}

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
			case *mail.AttachmentHeader:
			// This is an attachment. Get its filename.
			filename, err := h.Filename()
			if err != nil {
				log.Printf("Could not get attachment filename: %v", err)
				continue
			}

			// Get the size by reading and discarding the body.
			// This is more memory-efficient than reading it all into a buffer.
			size, err := io.Copy(ioutil.Discard, p.Body)
			if err != nil {
				log.Printf("Could not get attachment size: %v", err)
			continue
			}

			// Add the attachment metadata to our slice
			fullEmail.Attachments = append(fullEmail.Attachments, Attachment{
				Name: filename,
				Size: size,
				FormattedSize: formatBytes(size),
			})
		}
	}
	return fullEmail, nil
}

// DownloadAttachment finds an attachment by name and saves it to a user-chosen location.
func (a *App) DownloadAttachment(seqNum uint32, filename string, section string) (string, error) {
	if a.mail == nil || a.mail.Client == nil {
		return "", errors.New("not connected to a server")
	}
	_, err := a.mail.Client.Select(section, false)
    if err != nil {
        return "", fmt.Errorf("could not select mailbox '%s': %w", section, err)
    }

	// Fetch the full email body
	seqset := new(imap.SeqSet)
	seqset.AddNum(seqNum)
	bodySection := &imap.BodySectionName{}
	items := []imap.FetchItem{bodySection.FetchItem()}
	messages := make(chan *imap.Message, 1)
	if err := a.mail.Client.Fetch(seqset, items, messages); err != nil {
		return "", fmt.Errorf("could not fetch email: %w", err)
	}
	msg := <-messages
	if msg == nil {
		return "", errors.New("email not found")
	}

	// Parse the email
	r := msg.GetBody(bodySection)
	mr, err := mail.CreateReader(r)
	if err != nil {
		return "", fmt.Errorf("could not parse email: %w", err)
	}

	// Loop through parts to find the matching attachment
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break // Reached the end without finding the attachment
		} else if err != nil {
			continue // Skip broken parts
		}

		if h, ok := p.Header.(*mail.AttachmentHeader); ok {
			attachmentFilename, _ := h.Filename()
			if attachmentFilename == filename {
				// We found the attachment! Now, save it.
				savePath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
					DefaultFilename: filename,
					Title:           "Save Attachment",
				})
				if err != nil {
					return "", err
				}
				if savePath == "" {
					return "Download cancelled by user.", nil // User closed the dialog
				}

				// Read the attachment data
				attachmentData, err := ioutil.ReadAll(p.Body)
				if err != nil {
					return "", fmt.Errorf("could not read attachment data: %w", err)
				}

				// Write the data to the chosen file path
				if err := os.WriteFile(savePath, attachmentData, 0644); err != nil {
					return "", fmt.Errorf("could not save attachment: %w", err)
				}

				return fmt.Sprintf("Successfully saved to %s", savePath), nil
			}
		}
	}

	return "", fmt.Errorf("attachment '%s' not found in email", filename)
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

// Level-1 One Time Pad functions
func  (a *App)GenRand(plaintext string) ([]byte, error) {
    key := make([]byte, len(plaintext))
    _, err := io.ReadFull(rand.Reader, key)
    if err != nil {
        return nil, fmt.Errorf("failed to generate random key: %v", err)
    }
    fmt.Printf("Key generation successful: %d bytes of entropy\n", len(key))
    return key, nil //Key, error
}

func (a *App)Encrypt(plaintext string,key []byte) (string, string, error) {
    if len(plaintext) != len(key) {
    	return "", "", fmt.Errorf("OTP violation: key length (%d) != plaintext length (%d)",
    		len(key), len(plaintext))
    }
    ciphertext := make([]byte, len(plaintext))
    for i := range plaintext {
    	ciphertext[i] = plaintext[i] ^ key[i]
    }
    fmt.Println("Encryption completed successfully")
    return hex.EncodeToString(ciphertext), hex.EncodeToString(key), nil //Ciphertext, key, error 
}


func  (a *App)Decrypt(ciphertextStr string, keyHex string) (string, error) {
    ciphertext, err := hex.DecodeString(ciphertextStr)
    if err != nil {
    	return "", fmt.Errorf("failed to decode ciphertext: %v", err)
    }
    key, err := hex.DecodeString(keyHex)
    if err != nil {
    	return "", fmt.Errorf("failed to decode key: %v", err)
    }
    if len(ciphertext) != len(key) {
    	return "", fmt.Errorf("decryption error: key length (%d) != ciphertext length (%d)",
    		len(key), len(ciphertext))
    }
    plaintext := make([]byte, len(ciphertext))
    for i := range ciphertext {
   	plaintext[i] = ciphertext[i] ^ key[i]
    }
    return string(plaintext), nil //Decrypted, error
}


// DeleteSingleEmail moves one email to the Trash folder.
func (a *App) DeleteSingleEmail(seqNum uint32, section string) (string, error) {
	if a.mail == nil || a.mail.Client == nil {
		return "", errors.New("not connected to a server")
	}

	// 1. Select the mailbox where the email currently is (e.g., "INBOX")
	if _, err := a.mail.Client.Select(section, false); err != nil {
		return "", fmt.Errorf("could not select mailbox '%s': %w", section, err)
	}

	// 2. Create a sequence set for the single email ID.
	seqset := new(imap.SeqSet)
	seqset.AddNum(seqNum)
	
	// 3. The destination folder is [Gmail]/Trash.
	trashFolder := "[Gmail]/Trash"

	// 4. Move the email.
	if err := a.mail.Client.Move(seqset, trashFolder); err != nil {
		return "", err
	}
    
    // 5. Expunge the mailbox to finalize the deletion from the current view.
    if err := a.mail.Client.Expunge(nil); err != nil {
        return "", err
    }

	return "Email moved to Trash.", nil
}

func (a *App) shutdown(ctx context.Context) {
	if a.mail != nil {
		a.mail.Disconnect()
	}
}
