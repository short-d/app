package email

var _ Sender = (*SenderStub)(nil)

type SenderStub struct {
	sendError error
	sentEmail Email
}

func (s *SenderStub) SendEmail(email Email) error {
	if s.sendError != nil {
		return s.sendError
	}

	s.sentEmail = email
	return nil
}

func (s SenderStub) GetSentEmail() Email {
	return s.sentEmail
}

func (s *SenderStub) SetSendError(err error) {
	s.sendError = err
}

func NewSenderStub() SenderStub {
	return SenderStub{}
}
