package main

import (
	"log"
	"net/smtp"
	"os"

	"github.com/pullemax/mail-sender/mail"
	"github.com/pullemax/mail-sender/struts"
)

func main() {
	args := os.Args
	var auth smtp.Auth
	var msg []byte

	var smtpV = struts.Smtp{}

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-host":
			smtpV.SetHost(args[i+1])
		case "-port":
			smtpV.SetPort(args[i+1])
		case "-user":
			smtpV.SetUser(args[i+1])
		case "-password":
			smtpV.SetPassword(args[i+1])
		case "-from":
			smtpV.SetFrom(args[i+1])
		case "-image":
			smtpV.SetImage(true)
			smtpV.SetPathImage(args[i+1])
		case "-document":
			smtpV.SetDocument(true)
		}
	}

	if smtpV.GetUser() != "" && smtpV.GetPassword() != "" {
		auth = smtp.PlainAuth("", smtpV.GetUser(), smtpV.GetPassword(), smtpV.GetHost())
	}

	template := mail.ReadTemplate("./resources/template/template.html")

	m := struts.Mail{
		From:    smtpV.GetFrom(),
		To:      "pullemax@gmail.com",
		Subject: "Test de envío",
	}

	m.SetBody(template)

	if smtpV.GetImage() {
		m.SetImage(mail.ReadImages(smtpV.GetPathImage()))
	}

	msg = mail.BuildMessage(m)
	err := smtp.SendMail(smtpV.GetHost()+":"+smtpV.GetPort(), auth, smtpV.GetFrom(), []string{"pullemax@gmail.com"}, msg)
	if err != nil {
		log.Fatalln(err)
	}
}
