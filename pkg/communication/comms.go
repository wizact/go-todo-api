package comms

import (
	"github.com/sendgrid/sendgrid-go"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Emailer interface {
	Send() (int, error)
}

type Communication struct {
	emailer Emailer
}

func NewCommunication(em Emailer) Communication {
	return Communication{emailer: em}
}

type SendGridEmailer struct {
	SendGridKey string
}

func NewSendGridEmailer() SendGridEmailer {
	return SendGridEmailer{}
}

func (e SendGridEmailer) Send() (int, error) {
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
