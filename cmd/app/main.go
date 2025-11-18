package main

import (
	"context"

	"github.com/pc-configurator/components/config"
	"github.com/pc-configurator/components/internal/app"
	"github.com/pc-configurator/components/pkg/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		logger.Fatal(err, "config.New")
	}

	logger.Init(cfg.Logger)

	ctx := context.Background()
	err = app.Run(ctx, cfg)
	if err != nil {
		logger.Error(err, "app.Run")
	}
}
