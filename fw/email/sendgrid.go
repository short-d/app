package email

import (
	"errors"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var _ Sender = (*SendGrid)(nil)

const contentTypeHTML = "text/html"

type SendGrid struct {
	apiKey string
}

func (s SendGrid) SendEmail(email Email) error {
	from := mail.NewEmail(email.FromName, email.FromAddress)
	to := mail.NewEmail(email.ToName, email.ToAddress)
	content := mail.Content{
		Type:  contentTypeHTML,
		Value: email.ContentHTML,
	}
	sendGridMail := mail.NewV3MailInit(from, email.Subject, to, &content)
	client := sendgrid.NewSendClient(s.apiKey)
	res, err := client.Send(sendGridMail)

	if err != nil {
		return err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return errors.New(http.StatusText(res.StatusCode))
	}
	return nil
}

func NewSendGrid(apiKey string) SendGrid {
	return SendGrid{apiKey: apiKey}
}
