package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	header  string
	message string
	Tag     string
}

var Nil = Error{}

var (
	TagBadRequest          = "BAD_REQUEST"
	TagInternalServerError = "INTERNAL_SERVER_ERROR"
)

func New(header, message, tag string) Error {
	return Error{header: header, message: message, Tag: tag}
}

func (e *Error) Error() string {
	return fmt.Sprintf("HEADER: %s\nMESSAGE: %s\nTAG: %s", e.header, e.message, e.Tag)
}

func (e *Error) Is(target error) bool {
	var t *Error
	if errors.As(target, &t) {
		return e.header == t.header && e.message == t.message && e.Tag == t.Tag
	}
	return false
}
