package main

import (
	"errors"
	"log"
	"strconv"
	"gopkg.in/gomail.v2"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)


type Mail struct {
	MailID string
	Password string
	Client *client.Client
}

func NewMail(user_mailid string, user_password string) (*Mail, error) {
	if user_mailid == "" {
		return nil, errors.New("MailID Cannot Be Empty")
	}
	if user_password == "" {
		return nil, errors.New("Password Cannot Be Empty")
	}
	return &Mail {
		MailID: user_mailid,
		Password: user_password,
	}, nil
}

func (m *Mail) Connect(addr string) error {
	if m.Client != nil {
		return errors.New("Already Connected, Please disconnect first")
	}

	c, err := client.DialTLS(addr, nil)
	if err != nil {
		return err
	}
	log.Println("Connection Established")

	m.Client = c

	log.Println("Logging IN")
	
	err = m.Client.Login(m.MailID, m.Password)
	if err != nil {
		m.Client.Logout()
		return err
	}
	log.Println("Logged IN Successfully")
	return nil
}

func (m *Mail) Disconnect() {
	if m.Client != nil {
		log.Println("Disconnecting....")
		m.Client.Logout()
		m.Client = nil
		log.Println("Disconnected Successfully")

	}
}

func (m *Mail) SendEmail(to, cc, bcc []string, subject string, body string, attachmentPaths []string) error {
	msg := gomail.NewMessage()
	
	msg.SetHeader("From", m.MailID)
	msg.SetHeader("To", to...)
	if len(cc) > 0 {
		msg.SetHeader("Cc", cc...)
	}
	if len(bcc) > 0 {
		msg.SetHeader("Bcc", bcc...)
	}
	msg.SetHeader("Subject", subject)

	// Message Body
	msg.SetBody("text/html", body)

	// Attach files to msg
	for _, path := range attachmentPaths {
		if path != "" {
			msg.Attach(path)
		}
	}


	// Configure Port and Host
	port, _ := strconv.Atoi("587")
	dialer := gomail.NewDialer("smtp.gmail.com", port, m.MailID, m.Password)

	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}


func (m *Mail) GetMessages(mailbox string, count uint32) ([]*imap.Message, error) {
	if m.Client == nil {
		return nil, errors.New("not connected")
	}

	mbox, err := m.Client.Select(mailbox, false)
	if err != nil {
		return nil, err
	}

	if mbox.Messages == 0 {
		return []*imap.Message{}, nil
	}

	from := uint32(1)
	if mbox.Messages > count {
		from = mbox.Messages - count + 1
	}
	to := mbox.Messages

	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messagesChan := make(chan *imap.Message, count)
	done := make(chan error, 1)

	go func() {
		done <- m.Client.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, imap.FetchFlags}, messagesChan)
	}()

	if err := <- done; err != nil {
		return nil, err
	}

	var messages []*imap.Message
	for msg := range messagesChan {
		messages = append(messages, msg)
	}

	return messages, nil
}
