package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pc-configurator/components/pkg/httpserver"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/postgres"
)

type App struct {
	Name    string `envconfig:"APP_NAME" required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App      App
	Logger   logger.Config
	HTTP     httpserver.Config
	Postgres postgres.Config
}

func New() (Config, error) {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		return config, logger.NewErrorWithPath("godotenv.Load", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, logger.NewErrorWithPath("envconfig.Process", err)
	}

	return config, nil
}
