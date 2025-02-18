package service

import "log"

type MemoryEmailClient struct {
}

func NewMemoryEmailClient() MemoryEmailClient {
	return MemoryEmailClient{}
}

func (e MemoryEmailClient) Send(to, toemail, subj, plainText, htmlText string) (int, error) {
	log.Println("Sending email to: ", toemail)

	return 200, nil
}

func (e MemoryEmailClient) SendUsingTemplate(to, toemail, subj, templateId string, tdata map[string]string) (int, error) {
	log.Println("Sending email to: ", toemail, " using template: ", templateId, " with data: ", tdata)

	return 200, nil
}
