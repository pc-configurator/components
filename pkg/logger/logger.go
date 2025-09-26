package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	AppName       string `envconfig:"APP_NAME"               required:"true"`
	AppVersion    string `envconfig:"APP_VERSION"            required:"true"`
	Level         string `envconfig:"LOGGER_LEVEL"           default:"error"`
	PrettyConsole bool   `envconfig:"LOGGER_PRETTY_CONSOLE"  default:"false"`
}

func Init(c Config) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	level, err := zerolog.ParseLevel(c.Level)
	if err != nil {
		zerolog.SetGlobalLevel(level)
	}

	log.Logger = log.With().
		Caller().
		Str("app_name", c.AppName).
		Str("app_version", c.AppVersion).
		Logger()

	if c.PrettyConsole == true {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	}

	log.Info().Msg("Logger initialized")
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Error(err error, msg string) {
	log.Error().Err(err).Msg(msg)
}

func Fatal(err error, msg string) {
	log.Fatal().Err(err).Msg(msg)
}
