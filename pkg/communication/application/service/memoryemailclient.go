package service

import "log"

type MemoryEmailClient struct {
}

func NewMemoryEmailClient() MemoryEmailClient {
	return MemoryEmailClient{}
}

func (e MemoryEmailClient) Send(to, toEmail, subject, plainText, htmlText string) error {
	log.Println("Sending email to: ", toEmail)

	return nil
}

func (e MemoryEmailClient) SendUsingTemplate(to, toEmail, subject, templateID string, templateData map[string]string) error {
	log.Println("Sending email to: ", toEmail, " using template: ", templateID, " with data: ", templateData)

	return nil
}
