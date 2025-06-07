package validator

import (
	"errors"
	"github.com/gookit/validate"
)

// IsValidationError checks if an error is of type ValidationError and extracts the validation errors.
//
// This function takes an error and determines if it is a specific validation error ([ValidationError]).
// If it is, returns the validation errors map contained in the error along with true.
// If it's not, returns nil and false.
//
// Parameters:
//   - err error: The error to be checked
//
// Returns:
//   - errorsMap [validate.Errors]: A map of validation errors if the error is of type ValidationError, nil otherwise
//   - isErr bool: true if the error is of type ValidationError, false otherwise
//
// Example:
//
//	if errorsMap, isValidationErr := validator.IsValidationError(err); isValidationErr {
//	    // Handle validation errors
//	    return HandleValidationErrors(errorsMap)
//	}
func IsValidationError(err error) (errorsMap validate.Errors, isErr bool) {
	var asValidationErr *ValidationError
	isErr = errors.As(err, &asValidationErr)
	if isErr {
		return asValidationErr.Errors, isErr
	}
	return nil, isErr
}
