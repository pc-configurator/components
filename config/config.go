package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pc-configurator/components/pkg/base_errors"
	"github.com/pc-configurator/components/pkg/http_server"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/postgres"
)

type App struct {
	ENV     string `envconfig:"APP_ENV"         required:"true"`
	Name    string `envconfig:"APP_NAME"        required:"true"`
	Version string `envconfig:"APP_VERSION"     required:"true"`
}

type Config struct {
	App      App
	Logger   logger.Config
	Postgres postgres.Config
	HTTP     http_server.Config
}

func New() (Config, error) {
	var config Config

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return config, base_errors.WithPath("godotenv.Load", errors.New("config path is empty"))
	}

	err := godotenv.Load(configPath)
	if err != nil {
		return config, base_errors.WithPath("godotenv.Load", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, base_errors.WithPath("envconfig.Process", err)
	}

	return config, nil
}
