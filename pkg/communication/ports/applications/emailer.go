package service

// Emailer is an interface for sending emails
type Emailer interface {
	Send(to, toemail, subj, plainText, htmlText string) (int, error)
	SendUsingTemplate(to, toemail, subj, templateId string, tdata map[string]string) (int, error)
}
