package main

import (
	"bytes"
	"fmt"
	"golang-email-api/env"
	"golang-email-api/mail"
	"html/template"
	"io"

	"gopkg.in/gomail.v2"
)

var epr *env.EnvPropReader

func init() {
	epr = env.NewEnvPropReader()
	epr.FileName = "app"
	epr.FileType = "yaml"
	epr.Location = "../Golang-Email-API/"
	epr.ReadEnv()

}

func main() {
	fmt.Println("Hello World")
	// sendEmail()
	// sendEmailHTML()
	sendEmailGoMailAPI()
}

func sendEmailGoMailAPI() {
	m := gomail.NewMessage()
	m.SetHeader("From", "andhikarizki00000@gmail.com")
	m.SetHeader("To", "andhikarizki00001@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "Hello!")

	s := gomail.SendFunc(func(from string, to []string, msg io.WriterTo) error {
		// Implements you email-sending function, for example by calling
		// an API, or running postfix, etc.
		fmt.Println("From:", from)
		fmt.Println("To:", to)
		return nil
	})

	if err := gomail.Send(s, m); err != nil {
		panic(err)
	}

}

func sendEmailHTML() {
	var htmlBody bytes.Buffer
	t, err := template.ParseFiles("../Golang-Email-API/html/test.html")

	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(&htmlBody, struct {
		Name      string
		OtherName string
	}{
		Name:      "DHIKA",
		OtherName: "Lomba Lari",
	})

	simpleMail := mail.SimpleMail{
		Sender:      "andhikarizki00000@gmail.com",
		Receiver:    []string{"alkyaby@gmail.com", "andhikarizki00001@gmail.com"},
		AppPassword: epr.EnvVariable["app_password"].(string),
		Host:        "smtp.gmail.com",
		MessageType: "html",
		Body:        htmlBody.String(),
		Subject:     "Test untuk yang pertama kali boleh kali",
	}

	err = simpleMail.SendMail()

	if err != nil {
		panic(err)
	}
}

func sendEmail() {
	// Receiver:    []string{"samuelnatanael9914@gmail.com", "ryanprn@gmail.com", "andhikarizki00001@gmail.com"},

	simpleMail := mail.SimpleMail{
		Sender:      "andhikarizki00000@gmail.com",
		Receiver:    []string{"alkyaby@gmail.com"},
		AppPassword: epr.EnvVariable["app_password"].(string),
		Host:        "smtp.gmail.com",
		MessageType: "text",
		Body:        "Hai Apakabar GUYS - dhika",
		Subject:     "Test untuk yang pertama kali boleh kali",
	}

	err := simpleMail.SendMail()

	if err != nil {
		panic(err)
	}
}
