package handler

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		write, err := w.Write([]byte("ok"))
		if err != nil {
			log.Warn().Err(err).Msgf("Failed to write health check response. Wrote %d bytes", write)
		}
	})
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
