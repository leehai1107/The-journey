package swagger

import (
	"strings"

	"github.com/leehai1107/The-journey/docs"
	"github.com/leehai1107/The-journey/internal/pkg/apiwrapper"
	"github.com/leehai1107/The-journey/internal/pkg/errors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type swagger struct {
}

func NewSwagger() *swagger {
	return &swagger{}
}

func (m *swagger) Register(gGroup gin.IRouter) {
	g := gGroup.Group("")
	{
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		g.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func (m *swagger) SwaggerHandler(isProduction bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isProduction {
			apiwrapper.Abort(c, &apiwrapper.Response{Error: errors.New("not allow to access")})
			return
		}
		docs.SwaggerInfo.Host = strings.ToLower(c.Request.Host)
		docs.SwaggerInfo.BasePath = "internal/api/v1"
		c.Next()
	}
}
