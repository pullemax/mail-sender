package struts

type Mail struct {
	From     string
	To       string
	Subject  string
	Body     string
	Image    []Document
	Document []Document
}

func (s *Mail) SetFrom(from string) {
	s.From = from
}

func (s *Mail) SetTo(to string) {
	s.To = to
}

func (s *Mail) SetSubject(subject string) {
	s.Subject = subject
}

func (s *Mail) SetBody(body string) {
	s.Body = body
}

func (s *Mail) SetImage(image []Document) {
	s.Image = image
}

func (s *Mail) SetDocument(document []Document) {
	s.Document = document
}

func (s *Mail) GetFrom() string {
	return s.From
}

func (s *Mail) GetTo() string {
	return s.To
}

func (s *Mail) GetSubject() string {
	return s.Subject
}

func (s *Mail) GetBody() string {
	return s.Body
}

func (s *Mail) GetImage() []Document {
	return s.Image
}

func (s *Mail) GetDocument() []Document {
	return s.Document
}
