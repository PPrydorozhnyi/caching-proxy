package config

import (
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/rs/zerolog/log"
)

type Conf struct {
	Server ConfServer
	Log    ConfLog
}

type ConfServer struct {
	Port int `env:"SERVER_PORT,default=8080"`

	Timeout struct {
		Read  time.Duration `env:"SERVER_TIMEOUT_READ,default=30s,strict"`
		Write time.Duration `env:"SERVER_TIMEOUT_WRITE,default=15s,strict"`
		Idle  time.Duration `env:"SERVER_TIMEOUT_IDLE,default=2m,strict"`
	}

	Debug bool `env:"SERVER_DEBUG,default=false"`
}

type ConfLog struct {
	Level      string `env:"LOG_LEVEL,default=info"`
	Structured bool   `env:"LOG_STRUCTURED,default=false"`
}

func New() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatal().Stack().Err(err).Msgf("Failed to decode: %s", err)
	}

	return &c
}
