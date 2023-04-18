package framework

import (
	"github.com/pkg/errors"
	"testing"
)

func Test_NewErrorWithErrorType(t *testing.T) {
	var code ErrorType = 10000000
	err := code.New("this is an error")
	if appError, ok := err.(AppError); ok {
		if appError.errorType != code {
			t.Fatal("code check failed")
		}
		if appError.originalError.Error() != "this is an error" {
			t.Fatal("error message error")
		}
	} else {
		t.Fatal("New not returns AppError")
	}
}

func Test_NewfErrorWithErrorType(t *testing.T) {
	var code ErrorType = 10000000
	err := code.Newf("this is an error: %d", 1)
	if appError, ok := err.(AppError); ok {
		if appError.errorType != code {
			t.Fatal("code check failed")
		}
		if appError.originalError.Error() != "this is an error: 1" {
			t.Fatal("error message error: 1")
		}
	} else {
		t.Fatal("New not returns AppError")
	}
}

func Test_WrapWithErrorType(t *testing.T) {
	var code ErrorType = 10000000
	internalErr := errors.New("this is an internal error")
	err := code.Wrap(internalErr, "this is wrap message")
	if appError, ok := err.(AppError); ok {
		if appError.errorType != code {
			t.Fatal("code check failed")
		}
		if appError.originalError.Error() != "this is wrap message: this is an internal error" {
			t.Fatal("error message error")
		}
		if err.Error() != "this is wrap message: this is an internal error" {
			t.Fatal("error message from err direct error")
		}
	} else {
		t.Fatal("New not returns AppError")
	}
}

func Test_NewErrorDirect(t *testing.T) {
	err := NewError("this is a NoType error")
	if appError, ok := err.(AppError); ok {
		if appError.errorType != NoType {
			t.Fatal("code check failed")
		}
		if appError.originalError.Error() != "this is a NoType error" {
			t.Fatal("error message error")
		}
	} else {
		t.Fatal("New not returns AppError")
	}
}

func Test_NewfErrorDirect(t *testing.T) {
	err := NewErrorf("this is a NoType error: %d", 1)
	if appError, ok := err.(AppError); ok {
		if appError.errorType != NoType {
			t.Fatal("code check failed")
		}
		if appError.originalError.Error() != "this is a NoType error: 1" {
			t.Fatal("error message error")
		}
	} else {
		t.Fatal("New not returns AppError")
	}
}

func Test_WrapErrorDirect(t *testing.T) {
	var code ErrorType = 10000000
	err := code.New("this is an error")
	err = Wrap(err, "wrap message")
	if appError, ok := err.(AppError); ok {
		if appError.errorType != code {
			t.Fatal("code check failed")
		}
		if appError.originalError.Error() != "wrap message: this is an error" {
			t.Fatal("error message error")
		}
	} else {
		t.Fatal("New not returns AppError")
	}

	err = errors.New("error")
	err = Wrap(err, "wrap message")
	if appError, ok := err.(AppError); ok {
		if appError.errorType != NoType {
			t.Fatal("code check failed")
		}
		if appError.originalError.Error() != "wrap message: error" {
			t.Fatal("error message error")
		}
	} else {
		t.Fatal("New not returns AppError")
	}
	if GetType(err) != NoType {
		t.Fatal("GetType error")
	}
}
