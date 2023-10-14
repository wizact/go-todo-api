package httpmodels

// AppError as app error container
type AppError struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// FriendlyError is sanitised error message sent back to the user
type FriendlyError struct {
	Message string `json:"message"`
}
