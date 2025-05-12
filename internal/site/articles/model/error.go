package model

import "github.com/alexfalkowski/sasha/internal/site/meta"

// NewError for site.
func NewError(code int, err error) *Error {
	return &Error{code: code, err: err}
}

// Error for site.
type Error struct {
	Info *meta.Info
	err  error
	code int
}

// Error satisfies the error interface.
func (e *Error) Error() string {
	return e.err.Error()
}

// Error satisfies the Coder interface.
func (e *Error) Code() int {
	return e.code
}
