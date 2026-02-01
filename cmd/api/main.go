package main

import (
	"caching-proxy/config"
	"caching-proxy/handler"
	"caching-proxy/logging"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {

	c := config.New()
	logging.ConfigureLogger(&c.Log)

	mux := http.NewServeMux()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      mux,
		ReadTimeout:  c.Server.Timeout.Read,
		WriteTimeout: c.Server.Timeout.Write,
		IdleTimeout:  c.Server.Timeout.Idle,
	}

	handler.RegisterHandlers(mux)

	go func() {
		log.Info().Msgf("Starting server on port %d", c.Server.Port)

		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to shutdown server")
	}

	log.Info().Msg("Server stopped")
}
