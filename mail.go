package main

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"log"
	"errors"
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
		done <- m.Client.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messagesChan)
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

func (m *Mail) PrintMessages(messages chan *imap.Message) error {
	if messages == nil { return errors.New("*imap.Message chan empty")
	}
	for msg := range messages {
		fmt.Println("================================")
		for _, addr := range msg.Envelope.From {
			fmt.Println("Personal Name: ", addr.PersonalName)
			fmt.Println("MailBox Name: ", addr.MailboxName)
			fmt.Println("Host Name: ", addr.HostName)
			fmt.Println()
		}
		fmt.Println("Subject: ", msg.Envelope.Subject)
		fmt.Println("Date: ", msg.Envelope.Date)
		fmt.Println("================================")
		fmt.Println()
	}

	return nil
}


