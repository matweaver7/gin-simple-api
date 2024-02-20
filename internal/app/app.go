package app

import (
	"context"
	"sync"

	"matweaver.com/simple-rest-api/config"
	"matweaver.com/simple-rest-api/internal/controllers"
	"matweaver.com/simple-rest-api/internal/repository"
	"matweaver.com/simple-rest-api/internal/services"
)

type App struct {
	repositories repository.Repositories
	services     services.Services
	controllers  controllers.Controllers
	cfg          config.Config
}

func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	app := &App{
		cfg: *cfg,
	}

	return app, nil
}

func (a *App) Run(ctx context.Context, wg *sync.WaitGroup) {

}
