package model

import "github.com/alexfalkowski/sasha/internal/site/meta"

// Error for site.
type Error struct {
	Info *meta.Info
	Err  error
}

// Error satisfies the error interface.
func (e *Error) Error() string {
	return e.Err.Error()
}
