package comms

import (
	"github.com/sendgrid/sendgrid-go"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Emailer interface {
	Send(to, toemail string) (int, error)
}

type Communication struct {
	emailer Emailer
}

func NewCommunication(em Emailer) Communication {
	return Communication{emailer: em}
}

func (c *Communication) Communicate(to, toemail string) (int, error) {
	return c.emailer.Send(to, toemail)
}

type SendGridEmailer struct {
	SendGridKey       string
	SendGridFromName  string
	SendGridFromEmail string
}

func NewSendGridEmailer() SendGridEmailer {
	return SendGridEmailer{}
}

func (e SendGridEmailer) Send(to, toemail string) (int, error) {
	f := mail.NewEmail(e.SendGridFromName, e.SendGridFromEmail)
	s := "Verify your account"
	t := mail.NewEmail(to, toemail)
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
