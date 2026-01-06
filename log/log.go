package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"os"
	"strings"
)

// todo improve with https://github.com/rs/zerolog
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
