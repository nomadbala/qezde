package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	header  string
	message string
}

var Nil = Error{}

func New(header, message string) Error {
	return Error{header: header, message: message}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.header, e.message)
}

func (e *Error) Is(target error) bool {
	var t *Error
	if errors.As(target, &t) {
		return e.header == t.header && e.message == t.message
	}
	return false
}
