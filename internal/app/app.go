package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pc-configurator/components/config"
	"github.com/pc-configurator/components/internal/adapter/postgres"
	"github.com/pc-configurator/components/internal/controller/http"
	"github.com/pc-configurator/components/internal/usecase"
	"github.com/pc-configurator/components/pkg/httpserver"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/router"
)

func Run(ctx context.Context, c config.Config) error {
	r := router.New()
	uc := usecase.New(postgres.New())
	http.ComponentsRouter(r, uc)

	httpServer := httpserver.New(r, c.HTTP)

	logger.Info("App started!")

	<-listenCloseSignals()

	logger.Info("App got signal to stop")

	httpServer.Close()

	logger.Info("App stopped!")

	return nil
}

func listenCloseSignals() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}
