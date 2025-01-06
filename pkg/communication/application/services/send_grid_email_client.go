package service

import (
	"github.com/sendgrid/sendgrid-go"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const SENDGRID_ENDPOINT = "/v3/mail/send"
const SENDGRID_HOST = "https://api.sendgrid.com"

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

func (e SendGridEmailClient) Send(to, toemail, subj, plainText, htmlText string) (int, error) {
	f := mail.NewEmail(e.sendGridFromName, e.sendGridFromEmail)
	s := subj
	t := mail.NewEmail(to, toemail)
	ptc := plainText
	htc := htmlText
	ms := mail.NewSingleEmail(f, s, t, ptc, htc)
	client := sendgrid.NewSendClient(e.sendGridKey)
	r, err := client.Send(ms)
	if err != nil {
		return 0, err
	} else {
		return r.StatusCode, nil
	}
}

func (e SendGridEmailClient) SendUsingTemplate(to, toemail, subj, templateId string, tdata map[string]string) (int, error) {
	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail(e.sendGridFromName, e.sendGridFromEmail))
	m.SetTemplateID(templateId)

	// create personalization
	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail(to, toemail),
	}
	p.AddTos(tos...)

	for k, v := range tdata {
		p.SetDynamicTemplateData(k, v)
	}

	m.AddPersonalizations(p)
	req := sendgrid.GetRequest(e.sendGridKey, SENDGRID_ENDPOINT, SENDGRID_HOST)
	req.Method = "POST"
	var b = mail.GetRequestBody(m)
	req.Body = b
	res, err := sendgrid.API(req)
	if err != nil {
		return 0, err
	} else {
		return res.StatusCode, nil
	}
}
