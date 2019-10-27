package docopt

import (
	"fmt"
)

type errorType int

const (
	errorUser errorType = iota
	errorLanguage
)

func (e errorType) String() string {
	switch e {
	case errorUser:
		return "errorUser"
	case errorLanguage:
		return "errorLanguage"
	}
	return ""
}

// UserError records an error with program arguments.
type UserError struct {
	msg   string
	Usage string
}

func (e UserError) Error() string {
	return e.msg
}
func newUserError(msg string, f ...interface{}) error {
	return &UserError{fmt.Sprintf(msg, f...), ""}
}

// LanguageError records an error with the doc string.
type LanguageError struct {
	msg string
}

func (e LanguageError) Error() string {
	return e.msg
}
func newLanguageError(msg string, f ...interface{}) error {
	return &LanguageError{fmt.Sprintf(msg, f...)}
}

var newError = fmt.Errorf
