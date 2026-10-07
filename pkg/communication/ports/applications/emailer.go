package service

// Emailer is an interface for sending emails
type Emailer interface {
	Send(to, toEmail, subject, plainText, htmlText string) error
	SendUsingTemplate(to, toEmail, subject, templateID string, templateData map[string]string) error
}
