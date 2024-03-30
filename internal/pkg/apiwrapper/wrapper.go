package apiwrapper

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/The-journey/internal/pkg/errors"
)

const (
	statusSuccess = 1
	statusFail    = 0
)

/*
*
Example success response:

	{
	    "status": 1,
	    "code": 1,
	    "message": "Thành công!",
	    "data": {
	        "station": {
	            "code": "HUB210804225118",
	            "name": "Sen Station GV"
	        }
	    },
	    "trace_id": ""
	}

Example failure response:

	{
	    "status": 0,
	    "code": -1,
	    "message": "Yêu cầu không hợp lệ!",
	    "error": {
	        "code": -1,
	        "message": "Yêu cầu không hợp lệ!"
	    },
		"result":{
			"data":{}
		}
	    "trace_id": ""
	}
*/
type Response struct {
	Error      error       `json:"error"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
}

type APIResponse struct {
	Status  int64            `json:"status"`
	Code    errors.ErrorType `json:"code"`
	Message string           `json:"message"`
	Data    interface{}      `json:"data,omitempty"`
	Error   interface{}      `json:"error,omitempty"`
}
type apiHandlerFn func(c *gin.Context) *Response

func Wrap(fn apiHandlerFn) gin.HandlerFunc {
	return func(c *gin.Context) {
		rsp := fn(c)
		if rsp != nil {
			apiRsp := transformAPIResponse(c, rsp)
			if rsp.StatusCode == 0 {
				c.JSON(http.StatusOK, apiRsp)
				return
			}
			c.JSON(rsp.StatusCode, apiRsp)
		}
	}
}

func Abort(c *gin.Context, rsp *Response) {
	apiRsp := transformAPIResponse(c, rsp)
	if rsp.StatusCode == 0 {
		c.AbortWithStatusJSON(http.StatusOK, apiRsp)
		return
	}
	c.AbortWithStatusJSON(rsp.StatusCode, apiRsp)
}

func AbortWithStatus(c *gin.Context, statusCode int, rsp *Response) {
	apiRsp := transformAPIResponse(c, rsp)
	c.AbortWithStatusJSON(statusCode, apiRsp)
}

func File(fn apiHandlerFn) gin.HandlerFunc {
	return func(c *gin.Context) {
		rsp := fn(c)
		if rsp.StatusCode == 0 {
			fileName := rsp.Data.(string)
			c.FileAttachment(fileName, fileName)
			os.Remove(fileName)
			return
		}
		c.JSON(rsp.StatusCode, "error")
	}
}

func transformAPIResponse(c *gin.Context, rsp *Response) *APIResponse {
	if rsp == nil {
		return nil
	}
	if errors.Is(rsp.Error, errors.Success) {
		apiRsp := &APIResponse{
			Status:  statusSuccess,
			Code:    errors.GetErrorType(rsp.Error),
			Message: errors.GetMessage(rsp.Error),
			Data:    rsp.Data,
		}
		return apiRsp
	}
	apiRsp := &APIResponse{
		Status:  statusFail,
		Code:    errors.GetErrorType(rsp.Error),
		Message: errors.GetMessage(rsp.Error),
		Error:   rsp.Error,
		Data:    rsp.Data,
	}
	_ = c.Error(rsp.Error)
	return apiRsp
}
