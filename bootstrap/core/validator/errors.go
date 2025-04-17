package validator

import (
	"net/http"

	"github.com/gookit/validate"
)

type (
	ValidationError struct {
		Errors interface{} `json:"errors,inline"`
	}
)

func (c *ValidationError) Error() string {
	switch he := c.Errors.(type) {
	case validate.Errors:
		return he.Error()
	default:
		return "request is not valid. Some field does not fulfill requirements"
	}
}

func (c *ValidationError) StatusCode() int {
	return http.StatusBadRequest
}
