package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pc-configurator/components/config"
	"github.com/pc-configurator/components/pkg/postgres"

	"github.com/pc-configurator/components/internal/adapter/postgres_entities"
	"github.com/pc-configurator/components/internal/controller/http"
	"github.com/pc-configurator/components/internal/usecase"
	"github.com/pc-configurator/components/pkg/httpserver"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/router"
)

func Run(ctx context.Context, c config.Config) error {
	// Postgres
	pgpool, err := postgres.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("pgx.New: %w", err)
	}

	// UseCase
	uc := usecase.New(postgres_entities.New(pgpool))

	// HTTP
	r := router.New()
	http.ComponentRouter(r, uc)

	httpServer := httpserver.New(r, c.HTTP)

	logger.Info("App started")

	<-listenCloseSignals()

	logger.Info("App got signal to stop")

	httpServer.Close()
	pgpool.Close()

	logger.Info("App stopped")

	return nil
}

func listenCloseSignals() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}
