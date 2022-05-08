package errors

import (
	"errors"
	"fmt"
)

type UtilsError struct {
	Op  string
	Err error
}

func NewUtilsError(op string, message string) *UtilsError {
	return &UtilsError{Op: op, Err: errors.New(message)}
}

func NewUtilsErrorWrap(op string, err error) *UtilsError {
	return &UtilsError{Op: op, Err: err}
}

func (e *UtilsError) Error() string {
	return fmt.Sprintf("%s : %s", e.Op, e.Err.Error())
}

func (e *UtilsError) Unwrap() error { return e.Err }
