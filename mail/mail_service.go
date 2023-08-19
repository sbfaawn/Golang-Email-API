package mail

import (
	"errors"
	"fmt"
	"net/smtp"
)

type SimpleMail struct {
	Sender      string
	AppPassword string
	Receiver    []string
	Host        string
	Body        string
	Subject     string
	MessageType string
}

func (mail *SimpleMail) SendMail() error {
	message := ""
	if mail.MessageType == "text" {
		message = fmt.Sprint("Subject: ", mail.Subject, "\n", mail.Body)
	} else if mail.MessageType == "html" {
		headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
		message = fmt.Sprint("Subject: ", mail.Subject, "\n", headers, "\n\n", mail.Body)
	} else {
		return errors.New("message type is not recognized")
	}

	auth := smtp.PlainAuth(
		"",
		mail.Sender,
		mail.AppPassword,
		mail.Host,
	)

	err := smtp.SendMail(
		mail.Host+":587",
		auth,
		mail.Sender,
		mail.Receiver,
		[]byte(message),
	)

	if err != nil {
		return err
	}

	return nil
}
