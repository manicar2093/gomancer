package validator

// NoValidatorWarning represents a type used to indicate that validation is not implemented in this context. It says you need to use
// core.Validate function instead
type NoValidatorWarning struct{}

func (c NoValidatorWarning) Validate(i any) error {
	panic("gomancer does not uses this validation way. Use core.Validate instead")
}
