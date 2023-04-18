package framework

import (
	"fmt"
	"github.com/pkg/errors"
)

// ErrorType is the type of error
type ErrorType uint

const (
	// NoType error
	NoType ErrorType = iota
)

const (
	CommonInternalErr ErrorType = 10000000
)

const (
	EventKeyBizCode = "biz.code"
	EventKeyErrMsg  = "err.msg"
)

type AppError struct {
	errorType     ErrorType
	originalError error
}

// New creates a new AppError
func (errorType ErrorType) New(msg string) error {
	return AppError{errorType: errorType, originalError: errors.New(msg)}
}

// Newf creates a new AppError with formatted message
func (errorType ErrorType) Newf(msg string, args ...interface{}) error {
	return AppError{errorType: errorType, originalError: fmt.Errorf(msg, args...)}
}

// Wrap creates a new wrapped error
func (errorType ErrorType) Wrap(err error, msg string) error {
	return errorType.Wrapf(err, msg)
}

// Wrapf creates a new wrapped error with formatted message
func (errorType ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return AppError{errorType: errorType, originalError: errors.Wrapf(err, msg, args...)}
}

// Error returns the mssage of a AppError
func (error AppError) Error() string {
	return error.originalError.Error()
}

// NewError creates a no type error
func NewError(msg string) error {
	return AppError{errorType: NoType, originalError: errors.New(msg)}
}

// NewErrorf creates a no type error with formatted message
func NewErrorf(msg string, args ...interface{}) error {
	return AppError{errorType: NoType, originalError: errors.New(fmt.Sprintf(msg, args...))}
}

// Wrap an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Cause gives the original error
func Cause(err error) error {
	return errors.Cause(err)
}

// Wrapf an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(AppError); ok {
		return AppError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
		}
	}

	return AppError{errorType: NoType, originalError: wrappedError}
}

// GetType returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(AppError); ok {
		return customErr.errorType
	}
	return NoType
}
