package errors

// unprocessableEntity represents an error when a request has bad parameters.
type unprocessableEntity struct {
	Err
}

// UnprocessableEntityf returns an error which satisfies IsUnprocessableEntity().
func UnprocessableEntityf(format string, args ...interface{}) error {
	return &unprocessableEntity{wrap(nil, format, "", args...)}
}

// NewUnprocessableEntity returns an error which wraps err that satisfies
// IsUnprocessableEntity().
func NewUnprocessableEntity(err error, msg string) error {
	return &unprocessableEntity{wrap(err, msg, "")}
}

// IsUnprocessableEntity reports whether err was created with UnprocessableEntityf() or
// UnprocessableEntity().
func IsUnprocessableEntity(err error) bool {
	err = Cause(err)
	_, ok := err.(*unprocessableEntity)
	return ok
}
