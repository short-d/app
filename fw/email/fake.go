package email

var _ Sender = (*SenderFake)(nil)

type SenderFake struct {
	sendError error
	sentEmail Email
}

func (s *SenderFake) SendEmail(email Email) error {
	if s.sendError != nil {
		return s.sendError
	}

	s.sentEmail = email
	return nil
}

func (s SenderFake) GetSentEmail() Email {
	return s.sentEmail
}

func (s *SenderFake) SetSendError(err error) {
	s.sendError = err
}

func NewSenderFake() SenderFake {
	return SenderFake{}
}
