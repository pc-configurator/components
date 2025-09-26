package main

import (
	"context"

	"github.com/pc-configurator/components/config"
	"github.com/pc-configurator/components/internal/app"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/validation"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		logger.Fatal(err, "config.New")
	}

	logger.Init(cfg.Logger)
	validation.Init()

	err = app.Run(ctx, cfg)
	if err != nil {
		logger.Fatal(err, "app.Run")
	}

	logger.Info("App stopped!")
}
