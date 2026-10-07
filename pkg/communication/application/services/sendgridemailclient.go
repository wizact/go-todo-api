package service

import (
	"errors"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var ErrSendGridRejected = errors.New("sendgrid rejected email")

type SendGridEmailClient struct {
	sendGridFromName  string
	sendGridFromEmail string
	send              func(*mail.SGMailV3) (int, error)
}

func NewSendGridEmailClient(sgKey, sgFromName, sgFromEmail string) SendGridEmailClient {
	client := sendgrid.NewSendClient(sgKey)
	return SendGridEmailClient{
		sendGridFromName:  sgFromName,
		sendGridFromEmail: sgFromEmail,
		send: func(message *mail.SGMailV3) (int, error) {
			response, err := client.Send(message)
			if err != nil {
				return 0, err
			}
			return response.StatusCode, nil
		},
	}
}

func (e SendGridEmailClient) Send(to, toEmail, subject, plainText, htmlText string) error {
	from := mail.NewEmail(e.sendGridFromName, e.sendGridFromEmail)
	recipient := mail.NewEmail(to, toEmail)
	message := mail.NewSingleEmail(from, subject, recipient, plainText, htmlText)
	return e.sendMessage(message)
}

func (e SendGridEmailClient) sendMessage(message *mail.SGMailV3) error {
	statusCode, err := e.send(message)
	if err != nil {
		return fmt.Errorf("send email with SendGrid: %w", err)
	}
	if statusCode < 200 || statusCode >= 300 {
		return fmt.Errorf("%w: HTTP status %d", ErrSendGridRejected, statusCode)
	}

	return nil
}

func (e SendGridEmailClient) SendUsingTemplate(to, toEmail, subject, templateID string, templateData map[string]string) error {
	message := mail.NewV3Mail()
	message.SetFrom(mail.NewEmail(e.sendGridFromName, e.sendGridFromEmail))
	message.SetTemplateID(templateID)

	personalization := mail.NewPersonalization()
	personalization.AddTos(mail.NewEmail(to, toEmail))
	for key, value := range templateData {
		personalization.SetDynamicTemplateData(key, value)
	}

	message.AddPersonalizations(personalization)
	return e.sendMessage(message)
}
