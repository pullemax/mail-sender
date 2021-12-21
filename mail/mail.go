package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
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
		attachImage(image, boundary, &msg)
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

func ReadImages(path string) []struts.Image {
	var images []struts.Image

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			f, err2 := ioutil.ReadFile(path + "/" + file.Name())
			if err2 != nil {
				log.Println(err)
			}
			b := make([]byte, base64.StdEncoding.EncodedLen(len(f)))
			base64.StdEncoding.Encode(b, f)
			img := struts.Image{
				Id:      file.Name(),
				Content: b,
			}
			images = append(images, img)
		}
	}

	return images
}

func attachImage(image struts.Image, boundary string, msg *bytes.Buffer) {

	contentId := strings.Split(image.Id, ".")[0]

	msg.WriteString("\r\n")
	msg.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	msg.WriteString(fmt.Sprintf("Content-Type: image/jpeg; name=\"%s\"\r\n", image.Id))
	msg.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", image.Id))
	msg.WriteString("Content-Transfer-Encoding: base64\r\n")
	msg.WriteString(fmt.Sprintf("X-Attachment-Id: %s\r\n", contentId))
	msg.WriteString(fmt.Sprintf("Content-ID: <%s>\r\n", contentId))
	msg.Write(image.Content)
	msg.WriteString("\r\n")
}
