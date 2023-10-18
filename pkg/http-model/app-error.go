package httpmodels

// AppError as app error container
type AppError struct {
	ErrorObject error  `json:"error"`
	Message     string `json:"message"`
	Code        int    `json:"code"`
}

func (a *AppError) Error() string {
	return a.ErrorObject.Error()
}

// FriendlyError is sanitised error message sent back to the user
type FriendlyError struct {
	Message string `json:"message"`
}
