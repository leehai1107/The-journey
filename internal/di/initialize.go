package apifx

import (
	"github.com/leehai1107/The-journey/internal/pkg/infra"
	"github.com/leehai1107/The-journey/internal/service/blog/delivery/http"
	"github.com/leehai1107/The-journey/internal/service/blog/repository"
	"github.com/leehai1107/The-journey/internal/service/blog/usecase"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Provide(
	provideRouter,
	provideDB,
	provideHandler,
	provideRepo,
	provideUsecase,
)

func provideRouter(handler http.IHandler) http.Router {
	return http.NewRouter(handler)
}

func provideDB() *gorm.DB {
	return infra.GetDB()
}

func provideHandler(usecase usecase.IExampleUsecase) http.IHandler {
	return http.NewHandler(usecase)
}

func provideRepo(db *gorm.DB) repository.IExampleRepo {
	return repository.NewExampleRepo(db)
}

func provideUsecase(repo repository.IExampleRepo) usecase.IExampleUsecase {
	return usecase.NewExampleUsecase(repo)
}
