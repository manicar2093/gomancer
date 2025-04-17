package apperrors

type HandleableError interface {
	error
	StatusCode() int
}
