package service

// Emailer is an interface for sending emails
type Emailer interface {
	Send(to, toemail, subj, plainText, htmlText string) (int, error)
}
