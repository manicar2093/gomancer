package tags

import "strings"

const validateTagName = "validate"

type (
	ValidateGenerable interface {
		Generate() string
	}
	Validate struct {
		validations []ValidateGenerable
	}
	ValidateRequired     struct{}
	ValidateRequiredUuid struct{}
)

func NewValidate(validations ...ValidateGenerable) Generable {
	return &Validate{
		validations: validations,
	}
}

func NewValidateRequired() ValidateGenerable {
	return ValidateRequired{}
}

func NewValidateRequiredUuid() ValidateGenerable {
	return ValidateRequiredUuid{}
}

// Generate creates all configured validations separating them with pipe | as gookit/validate package request
func (v Validate) Generate() (string, string) {
	var (
		validationsTotalLength = len(v.validations)
		sb                     = &strings.Builder{}
	)

	for index, validation := range v.validations {
		if index+1 != validationsTotalLength {
			sb.WriteString("|")
		}
		sb.WriteString(validation.Generate())
	}

	return validateTagName, sb.String()
}

func (v ValidateRequired) Generate() string {
	return "required"
}

func (v ValidateRequiredUuid) Generate() string {
	return "required_uuid"
}
