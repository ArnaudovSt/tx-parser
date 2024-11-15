package errors

import "errors"

var (
	ErrAlreadyExists      = errors.New("already exists")
	ErrServiceUnavailable = errors.New("service unavailable")
)

func IsAlreadyExists(err error) bool {
	return errors.Is(err, ErrAlreadyExists)
}

func IsUnavailable(err error) bool {
	return errors.Is(err, ErrServiceUnavailable)
}
