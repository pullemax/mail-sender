package struts

import (
	"bytes"
	"fmt"
	"strings"
)

type Document struct {
	Id          string
	Content     []byte
	ContentType string
}

func (doc *Document) attach(boundary string, msg *bytes.Buffer) {

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
