package models

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// NewResponseError Return ResponseError with message init
func NewResponseError(mes string) *ResponseError {
	return &ResponseError{
		Message: mes,
	}
}
