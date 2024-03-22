package comms

import (
	"github.com/sendgrid/sendgrid-go"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	SendGridKey string
}

func NewEmail() *Email {
	return &Email{}
}

func (e *Email) Send() (int, error) {
	f := mail.NewEmail("hi", "hi@example.com")
	s := "Email Subject"
	t := mail.NewEmail("to", "to@example.com")
	ptc := "Please verify your account"
	htc := "<p>Please verify your account</p>"
	ms := mail.NewSingleEmail(f, s, t, ptc, htc)
	client := sendgrid.NewSendClient(e.SendGridKey)
	r, err := client.Send(ms)
	if err != nil {
		return 0, err
	} else {
		return r.StatusCode, nil
	}

}
