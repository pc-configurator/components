package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/pc-configurator/components/config"
	"github.com/pc-configurator/components/pkg/base_errors"
	"github.com/pc-configurator/components/pkg/http_server"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/postgres"
	"github.com/pc-configurator/components/pkg/router"
)

type Dependencies struct {
	RouterHTTP *chi.Mux
	Postgres   *postgres.Pool
}

func Run(ctx context.Context, cfg config.Config) (err error) {
	var deps Dependencies

	deps.Postgres, err = postgres.New(ctx, cfg.Postgres)
	if err != nil {
		return base_errors.WithPath("postgres.New", err)
	}
	defer deps.Postgres.Close()

	deps.RouterHTTP = router.New()
	//uc := usecase.New()

	httpServer := http_server.New(deps.RouterHTTP, cfg.HTTP.Port)
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
