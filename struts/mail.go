package struts

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type Mail struct {
	From     string
	To       string
	Subject  string
	Body     string
	Image    []Document
	Document []Document
}

func (m *Mail) BuildMessage() []byte {
	var msg bytes.Buffer
	boundary := uuid.NewString()

	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString(fmt.Sprintf("From: %s\r\n", m.From))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", m.To))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", m.Subject))
	msg.WriteString(fmt.Sprintf("Content-Type: multipart/related; boundary=\"%s\"\r\n", boundary))
	msg.WriteString(fmt.Sprintf("--%s\r\n", boundary))

	msg.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n")

	msg.WriteString(m.Body + "\r\n")

	for _, image := range m.Image {
		image.attach(boundary, &msg)
	}

	for _, doc := range m.Document {
		doc.attach(boundary, &msg)
	}

	msg.WriteString(fmt.Sprintf("--%s--\r\n", boundary))

	return msg.Bytes()
}

func (m *Mail) ReadTemplate(path string) string {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return ""
	}
	return string(content)
}

func (m *Mail) ReadFiles(path string) []Document {
	var documents []Document

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			if !strings.HasPrefix(file.Name(), ".") {
				f, err2 := ioutil.ReadFile(path + "/" + file.Name())
				if err2 != nil {
					log.Println(err)
				}
				contentType := http.DetectContentType(f)
				b := make([]byte, base64.StdEncoding.EncodedLen(len(f)))
				base64.StdEncoding.Encode(b, f)
				doc := Document{
					Id:          file.Name(),
					Content:     b,
					ContentType: contentType,
				}
				documents = append(documents, doc)
			}
		}
	}

	return documents
}
