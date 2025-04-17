package apperrors

type MessagedError struct {
	Message string `json:"error"`
	Code    int    `json:"code"`
}

func NewMessagedError(message string, code int) *MessagedError {
	return &MessagedError{
		Message: message,
		Code:    code,
	}
}

func (c MessagedError) Error() string {
	return c.Message
}

func (c MessagedError) StatusCode() int {
	return c.Code
}
