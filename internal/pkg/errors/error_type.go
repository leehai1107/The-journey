package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type ErrorType int

// New with custom error and default message
func (e ErrorType) New() error {
	return &customError{
		Code:          e,
		Message:       errorMap[e],
		originalError: errors.New(errorMap[e]),
	}
}

// Newm with custom error message
func (e ErrorType) Newm(msg string) error {
	return &customError{
		Code:          e,
		Message:       msg,
		originalError: errors.New(msg),
	}
}

// Newf with custom error and format message
func (e ErrorType) Newf(format string, args ...interface{}) error {
	return &customError{
		Code:          e,
		Message:       fmt.Sprintf(format, args...),
		originalError: fmt.Errorf(format, args...),
	}
}

// Wrap creates a new wrapped error
func (e ErrorType) Wrap(err error) error {
	return e.Wrapf(err, "")
}

// Wrap creates a new wrapped error
func (e ErrorType) Report(err error) error {
	return &customError{
		Code:          e,
		Message:       errorMap[e],
		originalError: errors.New(errorMap[e]),
		CauseError:    CustomError(err),
		ErrContext:    getErrorContext(err),
	}
}

// Wrapf creates a new wrapped error with formatted message
func (e ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return &customError{
		Code:          e,
		Message:       fmt.Sprintf(msg, args...),
		originalError: errors.Wrapf(err, msg, args...),
		CauseError:    Wrapf(err, msg, args...),
	}
}

// Cause gives the original Message
func Cause(err error) error {
	if customErr, ok := AsCustomError(err); ok {
		return customErr.Cause()
	}
	return errors.Cause(err)
}

// New Error
func New(msg string) error {
	return Newf(msg)
}

// Newf Error with format
func Newf(msg string, args ...interface{}) error {
	return Unknown.Newf(msg, args...)
}

// Cause gives the original Message
func GetErrorType(err error) ErrorType {
	if customErr, ok := AsCustomError(err); ok {
		return customErr.Code
	}
	return Unknown
}

// Wrap an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

func GetMessage(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// Wrapf an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := AsCustomError(err); ok {
		return &customError{
			Code:          customErr.Code,
			Message:       fmt.Sprintf(msg, args...),
			originalError: wrappedError,
			CauseError:    err,
		}
	}

	return &customError{
		Code:          Unknown,
		Message:       fmt.Sprintf(msg, args...),
		originalError: wrappedError,
		CauseError:    err,
	}
}

func AsCustomError(err error) (*customError, bool) {
	var customErr *customError
	if ok := errors.As(err, &customErr); ok {
		return customErr, true
	}
	return nil, false
}

//nolint:errorlint
func Is(err error, errorType ErrorType) bool {
	if e, ok := AsCustomError(err); ok {
		return e.Code == errorType
	}
	return false
}

func CustomError(err error) error {
	if err == nil {
		return nil
	}
	if customErr, ok := AsCustomError(err); ok {
		return customErr
	}

	return &customError{
		Code:          Unknown,
		Message:       err.Error(),
		originalError: err,
	}
}
