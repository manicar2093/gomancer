package converters

func Must[T any](ret T, err error) T { //nolint:ireturn
	if err != nil {
		panic(err)
	}
	return ret
}

func ToPointer[T any](d T) *T {
	return &d
}
