package errors

import (
	"errors"
	"fmt"
)

type EngineError struct {
	Op  string
	Err error
}

func NewEngineError(op string, message string) *EngineError {
	return &EngineError{Op: op, Err: errors.New(message)}
}

func NewEngineErrorWrap(op string, err error) *EngineError {
	return &EngineError{Op: op, Err: err}
}

func (e *EngineError) Error() string {
	return fmt.Sprintf("%s : %s", e.Op, e.Err.Error())
}

func (e *EngineError) Unwrap() error { return e.Err }
