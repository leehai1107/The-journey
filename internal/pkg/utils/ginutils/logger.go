package ginutils

import (
	"bytes"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/The-journey/internal/pkg/logger"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	_, err := w.body.Write(b)
	if err != nil {
		return 0, err
	}

	return w.ResponseWriter.Write(b)
}

// MiddlewareLogger for log request and response
func Logger(skipPaths ...string) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(skipPaths); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range skipPaths {
			skip[path] = struct{}{}
		}
	}
	return func(c *gin.Context) {
		lg := logger.EnhanceWith(c)
		start := time.Now()
		path := c.FullPath()

		// Check if the path matches any of the skip paths
		skipLogging := false
		for skipPath := range skip {
			if strings.HasPrefix(path, skipPath) {
				skipLogging = true
				break
			}
		}

		if !skipLogging {
			blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			c.Writer = blw
			c.Next()

			reqDump, err := httputil.DumpRequest(c.Request, true)
			if err != nil {
				return
			}
			reqDumpStr := string(reqDump)
			lg.Infow("--) [LOGGER] API middleware logger",
				"request", reqDumpStr,
				"response", blw.body.String(),
				"latency.ms", time.Since(start).Milliseconds(),
				"status", c.Writer.Status(),
				"ip", c.ClientIP())
		} else {
			// If skipping logging, just proceed to the next middleware
			c.Next()
		}
	}
}
