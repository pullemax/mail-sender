package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/pullemax/mail-sender/struts"
)

func BuildMessage(m struts.Mail) []byte {
	var msg bytes.Buffer
	boundary := uuid.NewString()

	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString(fmt.Sprintf("From: %s\r\n", m.GetFrom()))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", m.GetTo()))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", m.GetSubject()))
	msg.WriteString(fmt.Sprintf("Content-Type: multipart/related; boundary=\"%s\"\r\n", boundary))
	msg.WriteString(fmt.Sprintf("--%s\r\n", boundary))

	msg.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n")

	msg.WriteString(m.Body + "\r\n")

	for _, image := range m.GetImage() {
		attachDoc(image, boundary, &msg)
	}

	for _, doc := range m.GetDocument() {
		attachDoc(doc, boundary, &msg)
	}

	msg.WriteString(fmt.Sprintf("--%s--\r\n", boundary))

	return msg.Bytes()
}

func ReadTemplate(path string) string {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println(err)
		return ""
	}
	return string(content)
}

func ReadFiles(path string) []struts.Document {
	var documents []struts.Document

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
				doc := struts.Document{
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

func attachDoc(doc struts.Document, boundary string, msg *bytes.Buffer) {

	contentId := strings.Split(doc.Id, ".")[0]

	msg.WriteString("\r\n")
	msg.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	msg.WriteString(fmt.Sprintf("Content-Type: %s; name=\"%s\"\r\n", doc.ContentType, doc.Id))
	msg.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", doc.Id))
	msg.WriteString("Content-Transfer-Encoding: base64\r\n")
	msg.WriteString(fmt.Sprintf("X-Attachment-Id: %s\r\n", contentId))
	msg.WriteString(fmt.Sprintf("Content-ID: <%s>\r\n", contentId))
	msg.Write(doc.Content)
	msg.WriteString("\r\n")
}
