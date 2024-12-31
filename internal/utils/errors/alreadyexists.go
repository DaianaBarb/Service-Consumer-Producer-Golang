package errors

// alreadyExists represents and error when something already exists.
type alreadyExists struct {
	Err
}

// AlreadyExistsf returns an error which satisfies IsAlreadyExists().
func AlreadyExistsf(format string, args ...interface{}) error {
	sufix := ""
	if format == "" {
		sufix = " already exists"
	}
	return &alreadyExists{wrap(nil, format, sufix, args...)}
}

// NewAlreadyExists returns an error which wraps err and satisfies
// IsAlreadyExists().
func NewAlreadyExists(err error, msg string) error {
	return &alreadyExists{wrap(err, msg, "")}
}

// IsAlreadyExists reports whether the error was created with
// AlreadyExistsf() or NewAlreadyExists().
func IsAlreadyExists(err error) bool {
	err = Cause(err)
	_, ok := err.(*alreadyExists)
	return ok
}
