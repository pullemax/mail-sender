package main

import (
	"bytes"
	"html/template"
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
	var templateContent string
	var commonImages []struts.Document
	var commonDocuments []struts.Document

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
		case "-subject":
			smtpV.SetSubject(args[i+1])
		case "-image":
			smtpV.SetImage(true)
			smtpV.SetPathImage(args[i+1])
		case "-document":
			smtpV.SetDocument(true)
			smtpV.SetPathDocument(args[i+1])
		case "-template":
			smtpV.SetTemplate(args[i+1])
		case "-recipients":
			smtpV.SetRecipients(args[i+1])
		}
	}

	if smtpV.GetHost() == "" || smtpV.GetPort() == "" {
		log.Fatalln("SMTP host and port not found. You need to use -host and -port params")
	}

	if smtpV.GetUser() != "" && smtpV.GetPassword() != "" {
		auth = smtp.PlainAuth("", smtpV.GetUser(), smtpV.GetPassword(), smtpV.GetHost())
	}

	if smtpV.GetTemplate() != "" {
		templateContent = mail.ReadTemplate(smtpV.GetTemplate())
	} else {
		log.Fatalln("Not content to send. Use -template param and indicate the html or plain text file path")
	}

	recipients := mail.GetRecipients(smtpV.GetRecipients())

	if smtpV.GetImage() {
		commonImages = mail.ReadFiles(smtpV.GetPathImage())
	}

	if smtpV.GetDocument() {
		commonDocuments = mail.ReadFiles(smtpV.GetPathDocument())
	}

	for _, recipient := range recipients {
		var tempTemplate bytes.Buffer

		m := struts.Mail{
			From:    smtpV.GetFrom(),
			To:      recipient.Email,
			Subject: smtpV.GetSubject(),
		}

		t, errT := template.New("emailTemplate").Parse(templateContent)

		if errT != nil {
			log.Println(errT)
		}

		t.Execute(&tempTemplate, &recipient)

		m.SetBody(tempTemplate.String())
		m.SetImage(commonImages)
		m.SetDocument(commonDocuments)

		msg = mail.BuildMessage(m)
		err := smtp.SendMail(smtpV.GetHost()+":"+smtpV.GetPort(), auth, smtpV.GetFrom(), []string{recipient.Email}, msg)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
