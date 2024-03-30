package apiwrapper

import "github.com/leehai1107/The-journey/internal/pkg/errors"

func SuccessResponse() *Response {
	return &Response{
		Error: errors.Success.New(),
	}
}

func SuccessWithDataResponse(data interface{}) *Response {
	return &Response{
		Error: errors.Success.New(),
		Data:  data,
	}
}

func FailureResponse(err error) *Response {
	return &Response{
		Error: err,
	}
}
func FailureWithDataResponse(err error, data interface{}) *Response {
	return &Response{
		Error: err,
		Data:  data,
	}
}
