package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pc-configurator/components/config"
	"github.com/pc-configurator/components/internal/controller/http_router"
	"github.com/pc-configurator/components/internal/usecase"
	"github.com/pc-configurator/components/pkg/base_errors"
	"github.com/pc-configurator/components/pkg/http_server"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/router"
	"github.com/pc-configurator/components/pkg/validation"

	adapterPostgres "github.com/pc-configurator/components/internal/adapters/postgres"
	commonPostgres "github.com/pc-configurator/components/pkg/postgres"
)

func Run(ctx context.Context, cfg config.Config) error {
	pg, err := commonPostgres.New(ctx, cfg.Postgres)
	if err != nil {
		return base_errors.WithPath("postgres.New", err)
	}
	defer pg.Close()

	r := router.New()
	uc := usecase.New(adapterPostgres.New(pg))

	http_router.ComponentRouter(r, uc)

	validation.Init()

	httpServer := http_server.New(r, cfg.HTTP.Port)
	defer httpServer.Close()

	waiting(httpServer)

	return nil
}

func waiting(httpServer *http_server.Server) {
	logger.Info("App started")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	select {
	case i := <-wait:
		logger.Info("App got signal: " + i.String())
	case err := <-httpServer.Notify():
		logger.Error(err, "App got notify: httpServer.Notify")
	}

	logger.Info("App is stopping")
}
