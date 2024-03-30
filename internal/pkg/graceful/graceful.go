package graceful

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/The-journey/internal/pkg/logger"
)

const (
	TimeOutDefault  = 10 * time.Second
	DefaultWaitTime = 10 * time.Second
)

type Service interface {
	Register(g *gin.Engine)
	StartServer(handler http.Handler, port string)
	Close()
}

type service struct {
	currentStatus int
	waitTime      time.Duration
	timeout       time.Duration
	server        http.Server
}

func (s *service) Register(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "GREEN")
	})
}

func (s *service) StartServer(handler http.Handler, port string) {
	s.server = http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Errorf("failed to listen and serve from server: %v", err)
	}
}

func (s *service) stopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		logger.Errorw("server shutdown error", "error", err)
		return
	}
	logger.Info("stop server success")
}

func (s *service) Close() {
	logger.Info("set ping status to 503")
	s.currentStatus = http.StatusServiceUnavailable
	time.Sleep(s.waitTime)
	s.stopServer()
	logger.Info("server exited...")
}

func (s *service) SignalStop() {
	logger.Info("set ping status to 503")
	s.currentStatus = http.StatusServiceUnavailable
	time.Sleep(s.waitTime)
}

func NewService(opts ...Option) Service {
	o := &opt{waitTime: DefaultWaitTime, stopTimeout: TimeOutDefault}
	for _, opt := range opts {
		opt.apply(o)
	}
	return &service{
		currentStatus: http.StatusOK,
		waitTime:      o.waitTime,
		timeout:       o.stopTimeout,
	}
}
