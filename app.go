package main

import (
	"os"
	"log"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"errors"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
	mail *Mail
}

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

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) FetchEmails() ([]Email, error) {
	if a.mail == nil {
		return nil, errors.New("not connected to the mail server")
	}
	messages, err := a.mail.GetMessages("INBOX", 10)
	if err != nil {
		log.Fatal(err)
	}


	var emails []Email
	for _, msg := range messages {
		if msg.Envelope == nil {
			continue
		}

		from := ""
		if len(msg.Envelope.From) > 0 {
			addr := msg.Envelope.From[0]
			from = fmt.Sprintf("%s <%s@%s>",addr.PersonalName, addr.MailboxName, addr.HostName)
		}

		email := Email{
			From: from,
			Subject: msg.Envelope.Subject,
			Date: msg.Envelope.Date.Format("02 Jan 2006 15:04"),
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

func (a *App) SendEmail(to, cc, bcc string, subject string, body string) (string, error) {
	if a.mail == nil {
		return "",errors.New("Connect to a server first")
	}

	toList := cleanEmailList(to)
	ccList := cleanEmailList(cc)
	bccList := cleanEmailList(bcc)


	err := a.mail.SendEmail(toList, ccList, bccList, subject, body)
	if err != nil {
		return "", err
	}
	return "Sent Email Successfully", nil
	
}


func (a *App) shutdown(ctx context.Context) {
	if a.mail != nil {
		a.mail.Disconnect()
	}
}
