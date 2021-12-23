package main

import (
	"bytes"
	"flag"
	"html/template"
	"log"
	"net/smtp"

	"github.com/pullemax/mail-sender/mail"
	"github.com/pullemax/mail-sender/struts"
)

func main() {
	var auth smtp.Auth
	var msg []byte
	var templateContent string
	var commonImages []struts.Document
	var commonDocuments []struts.Document
	var smtpV = struts.Smtp{}

	flag.StringVar(&smtpV.Host, "host", "", "SMTP host")
	flag.StringVar(&smtpV.Port, "port", "", "SMTP port")
	flag.StringVar(&smtpV.User, "user", "", "SMTP auth user")
	flag.StringVar(&smtpV.Password, "password", "", "SMTP auth password")
	flag.StringVar(&smtpV.From, "from", "", "From email user")
	flag.StringVar(&smtpV.Subject, "subject", "", "Email Subject")
	flag.StringVar(&smtpV.PathImage, "image", "", "Image directory that the program will use to attach the images to the email")
	flag.StringVar(&smtpV.PathDocument, "document", "", "Document directory that the program will use to attach the docuents to the email")
	flag.StringVar(&smtpV.Template, "template", "", "Email template to send")
	flag.StringVar(&smtpV.Recipients, "recipients", "", "File with the recipients of the email")
	flag.Parse()

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

	if smtpV.GetPathImage() != "" {
		commonImages = mail.ReadFiles(smtpV.GetPathImage())
	}

	if smtpV.GetPathDocument() != "" {
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
