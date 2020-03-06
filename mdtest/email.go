package mdtest

import "github.com/short-d/app/fw"

var _ fw.EmailSender = (*EmailSenderFake)(nil)

type EmailSenderFake struct {
	sendError error
	sentEmail fw.Email
}

func (e *EmailSenderFake) SendEmail(email fw.Email) error {
	if e.sendError != nil {
		return e.sendError
	}

	e.sentEmail = email
	return nil
}

func (e EmailSenderFake) GetSentEmail() fw.Email {
	return e.sentEmail
}

func (e *EmailSenderFake) SetSendError(err error) {
	e.sendError = err
}

func NewEmailSenderFake() EmailSenderFake {
	return EmailSenderFake{}
}
