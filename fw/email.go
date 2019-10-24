package fw

type Email struct {
	FromName    string
	FromAddress string
	ToName      string
	ToAddress   string
	Subject     string
	ContentHTML string
}

type EmailSender interface {
	SendEmail(email Email) error
}
