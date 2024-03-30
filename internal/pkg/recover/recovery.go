package recover

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/The-journey/internal/pkg/apiwrapper"
	"github.com/leehai1107/The-journey/internal/pkg/errors"
	"github.com/leehai1107/The-journey/internal/pkg/logger"
)

func RPanic(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logger.EnhanceWith(c.Request.Context()).Errorf("method %v, path %v, err %v, trace %v",
				c.Request.Method,
				c.Request.URL.EscapedPath(),
				err,
				string(debug.Stack()))

			apiwrapper.Abort(c, &apiwrapper.Response{Error: errors.InternalServerError.New()})
		}
	}()

	c.Next()
}
