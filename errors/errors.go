package errors

import "errors"

var (
	ErrUnexpectedStatus  = errors.New("unexpected status code received from API")
	ErrUnsupportedMethod = errors.New("unsupported HTTP method")
)
