package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/The-journey/internal/pkg/logger"
)

type Router interface {
	Register(routerGroup gin.IRouter)
}

type routerIml struct {
	handler IHandler
}

func NewRouter(
	handler IHandler,
) Router {
	return &routerIml{
		handler: handler,
	}
}

func (p *routerIml) Register(r gin.IRouter) {
	lg := logger.EnhanceWith(context.Background())
	lg.Infow("RegisterRouterStart!")
	//routes for apis
	api := r.Group("api/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"now": time.Now(),
			})
		})
	}

	//routes for services
	svc := r.Group("")
	{
		svc.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"now": time.Now(),
			})
		})
	}
}
