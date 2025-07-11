package apperrors

import (
	"fmt"
	"github.com/gookit/validate"

	"github.com/manicar2093/gomancer/bootstrap/core/validator"
	"net/http"

	"github.com/coditory/go-errors"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

type (
	BasicError struct {
		Code  int         `json:"code"`
		Error interface{} `json:"error"`
	}
	ValidationError struct {
		validate.Errors `json:"errors"`
		Code            int `json:"code"`
	}
)

func HandlerWEcho(err error, ctx echo.Context) {
	code, body := handleErrorType(err) //nolint:varnamelen
	fields := log.Fields{"response": body}
	stack, ok := hasStackTrace(err)
	if ok {
		fields["stack"] = stack
	}
	log.WithFields(fields).Error(err)
	ctx.JSON(code, body) //nolint:errcheck
}

func handleErrorType(err error) (int, interface{}) {
	var (
		code    = http.StatusInternalServerError
		iterErr = err
	)

	for iterErr != nil {
		switch hErr := iterErr.(type) {
		case *validator.ValidationError:
			return http.StatusBadRequest, ValidationError{
				Code:   http.StatusBadRequest,
				Errors: hErr.Errors,
			}
		case HandleableError:
			return hErr.StatusCode(), BasicError{
				Code:  hErr.StatusCode(),
				Error: hErr.Error(),
			}
		case *echo.HTTPError:
			return hErr.Code, BasicError{
				Code:  hErr.Code,
				Error: hErr.Error(),
			}
		}
		iterErr = errors.Unwrap(iterErr)
		continue
	}
	return code, BasicError{
		Code:  code,
		Error: err.Error(),
	}
}

func hasStackTrace(err error) ([]string, bool) {
	var stack []string
	stacked, ok := err.(*errors.Error)
	if !ok {
		return nil, false
	}
	for _, d := range stacked.StackTrace() {
		stack = append(stack, fmt.Sprintf("%s:%d", d.RelFileName(), d.FileLine()))
	}
	return stack, true
}
