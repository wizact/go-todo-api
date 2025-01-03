package service

import (
	"github.com/sendgrid/sendgrid-go"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridEmailClient struct {
	sendGridKey       string
	sendGridFromName  string
	sendGridFromEmail string
}

func NewSendGridEmailClient(sgKey, sgFromName, sgFromEmail string) SendGridEmailClient {
	return SendGridEmailClient{
		sendGridKey:       sgKey,
		sendGridFromName:  sgFromName,
		sendGridFromEmail: sgFromEmail,
	}
}

func (e SendGridEmailClient) Send(to, toemail string) (int, error) {
	f := mail.NewEmail(e.sendGridFromName, e.sendGridFromEmail)
	s := "Verify your account"
	t := mail.NewEmail(to, toemail)
	ptc := "Please verify your account"
	htc := "<p>Please verify your account</p>"
	ms := mail.NewSingleEmail(f, s, t, ptc, htc)
	client := sendgrid.NewSendClient(e.sendGridKey)
	r, err := client.Send(ms)
	if err != nil {
		return 0, err
	} else {
		return r.StatusCode, nil
	}
}
