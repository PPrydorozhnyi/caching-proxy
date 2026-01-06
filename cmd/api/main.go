package main

import (
	"caching-proxy/config"
	"caching-proxy/logging"
	"errors"
	"fmt"
	"net/http"

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

	mux.HandleFunc("/", indexHandler)

	log.Info().Msgf("Starting server on port %d", c.Server.Port)

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
