package struts

type Smtp struct {
	host      string
	port      string
	user      string
	password  string
	from      string
	image     bool
	pathImage string
	document  bool
}

func (s *Smtp) SetHost(host string) {
	s.host = host
}

func (s *Smtp) SetPort(port string) {
	s.port = port
}

func (s *Smtp) SetUser(user string) {
	s.user = user
}

func (s *Smtp) SetPassword(password string) {
	s.password = password
}

func (s *Smtp) SetFrom(from string) {
	s.from = from
}

func (s *Smtp) SetImage(image bool) {
	s.image = image
}

func (s *Smtp) SetPathImage(pathImage string) {
	s.pathImage = pathImage
}

func (s *Smtp) SetDocument(document bool) {
	s.document = document
}

func (s *Smtp) GetHost() string {
	return s.host
}

func (s *Smtp) GetPort() string {
	return s.port
}

func (s *Smtp) GetUser() string {
	return s.user
}

func (s *Smtp) GetPassword() string {
	return s.password
}

func (s *Smtp) GetFrom() string {
	return s.from
}

func (s *Smtp) GetImage() bool {
	return s.image
}

func (s *Smtp) GetDocument() bool {
	return s.document
}

func (s *Smtp) GetPathImage() string {
	return s.pathImage
}
