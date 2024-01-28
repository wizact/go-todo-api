package httpservermodel

// AppError as app error container
type AppError struct {
	ErrorObject      error  `json:"error"`
	SanitisedMessage string `json:"message"`
	Code             int    `json:"code"`
}

func NewAppError(errorObject error, message string, code int) *AppError {
	return &AppError{ErrorObject: errorObject, SanitisedMessage: message, Code: code}
}

// Error returns the error string, or Message string in that order or priority.
func (a *AppError) Error() string {
	if a.ErrorObject == nil {
		return a.SanitisedMessage
	}

	return a.ErrorObject.Error()
}

// FriendlyError is sanitised error message sent back to the user
type FriendlyError struct {
	Message string `json:"message"`
}
