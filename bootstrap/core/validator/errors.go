package validator

import (
	"net/http"

	"github.com/gookit/validate"
)

type (
	ValidationError struct {
		validate.Errors
	}
)

func (c ValidationError) Error() string {
	return c.Errors.Error()
}

func (c ValidationError) StatusCode() int {
	return http.StatusBadRequest
}
