package validator

import (
	"github.com/google/uuid"
	"github.com/gookit/validate"
)

type GooKitValidator struct {
	StopOnError bool
}

const (
	uuidNotValidError    = "is not valid for UUID type"
	requiredErrorMessage = "needs to be on request"
	requiredUuidKey      = "required_uuid"
)

func NewGooKitValidator() *GooKitValidator {
	return &GooKitValidator{
		StopOnError: false,
	}
}

// ValidateStruct validates struct fulfill validate tags validations
//
// It adds a custom required validation:
//   - required_uuid
//
// This was created to validate uuid.UUID type from google UUID package.
func (c *GooKitValidator) ValidateStruct(toValidate any, scene ...string) error {
	v := c.GetConfiguredValidator(toValidate, scene...)
	if v.Validate(scene...) {
		return nil
	}
	return &ValidationError{Errors: v.Errors}
}

func (c *GooKitValidator) Validate(i interface{}) error {
	return c.ValidateStruct(i)
}

// GetConfiguredValidator returns validate.Validation instance with required_uuid validator
func (c *GooKitValidator) GetConfiguredValidator(toValidate any, scene ...string) *validate.Validation {
	validator := validate.Struct(toValidate, scene...) //nolint:varnamelen
	validator.StopOnError = c.StopOnError
	validator.AddMessages(map[string]string{
		"uuid":          uuidNotValidError,
		"isUUID":        uuidNotValidError,
		"required":      requiredErrorMessage,
		requiredUuidKey: requiredErrorMessage,
	})
	validator.AddValidator(requiredUuidKey, func(val any) bool {
		return val != uuid.Nil
	})
	return validator
}
