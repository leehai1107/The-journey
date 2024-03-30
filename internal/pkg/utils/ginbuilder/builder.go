package ginbuilder

import (
	"github.com/gin-gonic/gin"
	"github.com/leehai1107/The-journey/internal/pkg/utils/ginutils"
)

type builder struct {
	middlewares []gin.HandlerFunc
}

func BaseBuilder() *builder {
	return &builder{
		middlewares: []gin.HandlerFunc{
			gin.Recovery(),
		},
	}
}

func Default() *builder {
	return &builder{}
}

func (b *builder) WithBodyLogger(skipPaths ...string) *builder {
	b.middlewares = append(b.middlewares, ginutils.Logger(skipPaths...))
	return b
}

func (b *builder) Build() *gin.Engine {
	e := defaultGinEngine()
	e.Use(b.middlewares...)
	return e
}

func defaultGinEngine() *gin.Engine {
	e := gin.New()
	return e
}
