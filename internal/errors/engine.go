package errors

import (
	"errors"
	"fmt"
)

type EngineErrorType uint8

const (
	General EngineErrorType = iota
	AttackerFactory
	AttackerDead
	AttackerTrapped
	EndOfTheWorld
)

type EngineError struct {
	Op  EngineErrorType
	Err error
}

func NewEngineError(message string) *EngineError {
	return &EngineError{Op: General, Err: errors.New(message)}
}

func NewEngineErrorOp(op EngineErrorType) *EngineError {
	switch op {
	case EndOfTheWorld:
		return &EngineError{Op: op, Err: errors.New("it's the end of the world as we know it")}
	default:
		return &EngineError{Op: General, Err: errors.New("something unexpected happened")}
	}
}

func NewEngineErrorWrap(op EngineErrorType, err error) *EngineError {
	return &EngineError{Op: op, Err: err}
}

func (e *EngineError) Error() string {
	return fmt.Sprintf("%d : %s", e.Op, e.Err.Error())
}

func (e *EngineError) Unwrap() error { return e.Err }
