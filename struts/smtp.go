package struts

type Smtp struct {
	Host         string
	Port         string
	User         string
	Password     string
	From         string
	Subject      string
	PathImage    string
	PathDocument string
	Template     string
	Recipients   string
}

func (s *Smtp) SetHost(host string) {
	s.Host = host
}

func (s *Smtp) SetPort(port string) {
	s.Port = port
}

func (s *Smtp) SetUser(user string) {
	s.User = user
}

func (s *Smtp) SetPassword(password string) {
	s.Password = password
}

func (s *Smtp) SetFrom(from string) {
	s.From = from
}

func (s *Smtp) SetSubject(subject string) {
	s.Subject = subject
}

func (s *Smtp) SetPathImage(pathImage string) {
	s.PathImage = pathImage
}

func (s *Smtp) SetPathDocument(pathDocument string) {
	s.PathDocument = pathDocument
}

func (s *Smtp) SetTemplate(template string) {
	s.Template = template
}

func (s *Smtp) SetRecipients(recipients string) {
	s.Recipients = recipients
}

func (s *Smtp) GetHost() string {
	return s.Host
}

func (s *Smtp) GetPort() string {
	return s.Port
}

func (s *Smtp) GetUser() string {
	return s.User
}

func (s *Smtp) GetPassword() string {
	return s.Password
}

func (s *Smtp) GetFrom() string {
	return s.From
}

func (s *Smtp) GetSubject() string {
	return s.Subject
}

func (s *Smtp) GetPathImage() string {
	return s.PathImage
}

func (s *Smtp) GetPathDocument() string {
	return s.PathDocument
}

func (s *Smtp) GetTemplate() string {
	return s.Template
}

func (s *Smtp) GetRecipients() string {
	return s.Recipients
}
