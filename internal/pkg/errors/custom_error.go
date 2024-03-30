package errors

import "errors"

type customError struct {
	Code          ErrorType `json:"code"`
	Message       string    `json:"message"`
	StackTrace    string    `json:"stack_trace,omitempty"`
	CauseError    error     `json:"-"`
	originalError error
	ErrContext    errorContext `json:"error_context,omitempty"`
}

type errorContext map[string]string

func (e *customError) Error() string {
	return e.originalError.Error()
}

func (e *customError) Cause() error {
	if e.CauseError != nil {
		return e.CauseError
	}
	return e.originalError
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field, message string) error {
	var customErr *customError
	if ok := errors.As(err, &customErr); ok {
		if customErr.ErrContext == nil {
			customErr.ErrContext = make(map[string]string)
		}
		customErr.ErrContext[field] = message
		return customErr
	}
	return AddErrorContext(CustomError(err), field, message)
}

// GetErrorContext returns the error context
func getErrorContext(err error) errorContext {
	var customErr *customError
	if ok := errors.As(err, &customErr); ok && len(customErr.ErrContext) > 0 {
		return customErr.ErrContext
	}
	return errorContext{}
}
