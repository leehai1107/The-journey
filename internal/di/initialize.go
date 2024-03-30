package apifx

import (
	"github.com/leehai1107/The-journey/internal/pkg/infra"
	"github.com/leehai1107/The-journey/internal/service/blog/delivery/http"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Provide(
	provideRouter,
	provideDB,
	provideHandler,
)

func provideRouter(handler http.IHandler) http.Router {
	return http.NewRouter(handler)
}

func provideDB() *gorm.DB {
	return infra.GetDB()
}

func provideHandler() http.IHandler {
	return http.NewHandler()
}
