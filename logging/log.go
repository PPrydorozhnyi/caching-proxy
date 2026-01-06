package logging

import (
	"caching-proxy/config"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"os"
	"strings"
)

// check fo more options https://github.com/rs/zerolog
// todo add Integration with net/http
// todo add tracing hook
func ConfigureLogger(cfg *config.ConfLog) {

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	switch cfg.Level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	if !cfg.Structured {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}

}

func NewLogger(component string) zerolog.Logger {
	level := zerolog.GlobalLevel()

	if env := os.Getenv("LOG_LEVEL_" + strings.ToUpper(component)); env != "" {
		if l, err := zerolog.ParseLevel(env); err == nil {
			level = l
		}
	}

	return log.With().
		Str("component", component).
		Logger().
		Level(level)
}
