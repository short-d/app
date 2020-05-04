package email

type Email struct {
	FromName    string
	FromAddress string
	ToName      string
	ToAddress   string
	Subject     string
	ContentHTML string
}

type Sender interface {
	SendEmail(email Email) error
}
