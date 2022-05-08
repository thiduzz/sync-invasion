package errors

import (
	"errors"
	"fmt"
)

type CommandError struct {
	Parameter string
	Err       error
}

func NewCommandError(parameter string, message string) *CommandError {
	return &CommandError{Parameter: parameter, Err: errors.New(message)}
}

func (e *CommandError) Error() string {
	return fmt.Sprintf("%s : %s", e.Parameter, e.Err.Error())
}

func (e *CommandError) Unwrap() error { return e.Err }
