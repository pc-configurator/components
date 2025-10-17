package main

import (
	"context"

	"github.com/pc-configurator/components/config"
	"github.com/pc-configurator/components/internal/app"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	logger.Init(cfg.Logger)

	ctx := context.Background()
	err = app.Run(ctx, cfg)
	if err != nil {
		log.Error().Err(err).Msg("app.Run")
	}
}
