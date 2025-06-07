package validator

import (
	"github.com/google/uuid"
	"github.com/gookit/validate"
	"github.com/manicar2093/gomancer/bootstrap/cmd/service/translations/es"
)

const (
	requiredErrorMessage = "needs to be on request"
	requiredUuidKey      = "required_uuid"
)

func ConfiguredValidator(val *validate.Validation, lang string) {
	val.StopOnError = false
	val.AddValidator(requiredUuidKey, func(val any) bool {
		return val != uuid.Nil
	})
	if lang == "es" {
		es.Register(val)
	} else {
		val.AddMessages(map[string]string{
			requiredUuidKey: requiredErrorMessage,
		})
	}

}

func NewStructValidatorConfigured(data any, lang string, scene ...string) *validate.Validation {
	val := validate.Struct(data, scene...)
	ConfiguredValidator(val, lang)
	return val
}

func ValidateAndTransform(val *validate.Validation, scene ...string) error {
	if val.Validate(scene...) {
		return nil
	}
	return &ValidationError{Errors: val.Errors}
}
