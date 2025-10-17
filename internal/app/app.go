package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pc-configurator/components/config"
	"github.com/pc-configurator/components/internal/controller/http"
	"github.com/pc-configurator/components/internal/usecase"
	"github.com/pc-configurator/components/pkg/httpserver"
	"github.com/pc-configurator/components/pkg/router"
	"github.com/rs/zerolog/log"
)

func Run(ctx context.Context, c config.Config) error {
	r := router.New()
	uc := usecase.New()
	http.ComponentsRouter(r, uc)

	httpServer := httpserver.New(r, c.HTTP)

	log.Info().Msg("App started!")

	<-listenCloseSignals()

	log.Info().Msg("App got signal to stop")

	httpServer.Close()

	log.Info().Msg("App stopped!")

	return nil
}

func listenCloseSignals() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}
