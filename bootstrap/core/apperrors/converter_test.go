package apperrors_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/coditory/go-errors"
	"github.com/labstack/echo/v4"
	"github.com/manicar2093/gomancer/bootstrap/core/apperrors"
	"github.com/manicar2093/gomancer/bootstrap/core/validator"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("Converter", func() {
	var (
		echoInstance *echo.Echo
		req          *http.Request
		res          *httptest.ResponseRecorder
		echoCtx      echo.Context
	)

	BeforeEach(func() {
		echoInstance = echo.New()
		req, res = httptest.NewRequest(http.MethodGet, "/", nil), httptest.NewRecorder()
		echoCtx = echoInstance.NewContext(req, res)
	})

	Describe("HandlerWEcho", func() {
		When("receives a native error", func() {
			It("response as Internal Server Error with err as message", func() {
				var (
					err = fmt.Errorf("native error :(")
				)

				apperrors.HandlerWEcho(err, echoCtx)

				var mapRes map[string]interface{}
				Expect(json.NewDecoder(res.Body).Decode(&mapRes))
				Expect(mapRes).To(MatchKeys(IgnoreExtras, Keys{
					"error": Equal(err.Error()),
				}))
			})
		})
		When("error implements HandleableError", func() {
			It("response has status code and message from error", func() {
				var (
					err = &apperrors.MessagedError{
						Message: "error message",
						Code:    http.StatusAccepted,
					}
				)

				apperrors.HandlerWEcho(err, echoCtx)

				var mapRes map[string]interface{}
				Expect(json.NewDecoder(res.Body).Decode(&mapRes))
				Expect(mapRes).To(MatchKeys(IgnoreExtras, Keys{
					"error": Equal(err.Message),
					"code":  Equal(float64(err.Code)),
				}))
			})
		})
		When("error is wrapperd", func() {
			It("can handle it", func() {
				var (
					expectedMessage = "error message"
					expectedCode    = http.StatusAlreadyReported
					err             = errors.Wrap(&apperrors.MessagedError{
						Message: expectedMessage,
						Code:    expectedCode,
					})
				)

				apperrors.HandlerWEcho(err, echoCtx)

				var mapRes map[string]interface{}
				Expect(json.NewDecoder(res.Body).Decode(&mapRes))
				Expect(mapRes).To(MatchKeys(IgnoreExtras, Keys{
					"error": Equal(expectedMessage),
					"code":  Equal(float64(expectedCode)),
				}))
			})
		})

		When("error is validationError", func() {
			It("can handle it", func() {
				var (
					expectedCode = http.StatusBadRequest
					err          = &validator.ValidationError{
						Errors: map[string]interface{}{
							"name": "required",
						},
					}
				)

				apperrors.HandlerWEcho(err, echoCtx)

				var mapRes map[string]interface{}
				Expect(json.NewDecoder(res.Body).Decode(&mapRes))
				Expect(mapRes).To(MatchKeys(IgnoreExtras, Keys{
					"errors": MatchKeys(IgnoreExtras, Keys{
						"name": Equal("required"),
					}),
					"code": Equal(float64(expectedCode)),
				}))
			})
		})
	})

})
