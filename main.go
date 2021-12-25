package main

import (
	"bytes"
	"flag"
	"html/template"
	"log"
	"net/smtp"

	"github.com/pullemax/mail-sender/struts"
)

func main() {
	var auth smtp.Auth
	var templateContent string
	var commonImages []struts.Document
	var commonDocuments []struts.Document
	var smtpV = struts.Smtp{}
	var r struts.Recipient
	var m struts.Mail

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

	if smtpV.Host == "" || smtpV.Port == "" {
		log.Fatalln("SMTP host and port not found. You need to use -host and -port params")
	}

	if smtpV.User != "" && smtpV.Password != "" {
		auth = smtp.PlainAuth("", smtpV.User, smtpV.Password, smtpV.Host)
	}

	if smtpV.Template != "" {
		templateContent = m.ReadTemplate(smtpV.Template)
	} else {
		log.Fatalln("Not content to send. Use -template param and indicate the html or plain text file path")
	}

	recipients, err := r.GetRecipients(smtpV.Recipients)

	if err != nil {
		log.Fatal(err)
	}

	if smtpV.PathImage != "" {
		commonImages = m.ReadFiles(smtpV.PathImage)
	}

	if smtpV.PathDocument != "" {
		commonDocuments = m.ReadFiles(smtpV.PathDocument)
	}

	for _, recipient := range recipients {
		var tempTemplate bytes.Buffer

		m := struts.Mail{
			From:    smtpV.From,
			To:      recipient.Email,
			Subject: smtpV.Subject,
		}

		t, errT := template.New("emailTemplate").Parse(templateContent)

		if errT != nil {
			log.Println(errT)
		}

		t.Execute(&tempTemplate, &recipient)

		m.Body = tempTemplate.String()
		m.Image = commonImages
		m.Document = commonDocuments

		if err := smtp.SendMail(smtpV.Host+":"+smtpV.Port, auth, smtpV.From, []string{recipient.Email}, m.BuildMessage()); err != nil {
			log.Fatalln(err)
		}
	}
}
