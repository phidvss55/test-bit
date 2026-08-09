package api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	Config Config
}

type Config struct {
	Addr string
}

func (app *Application) Mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)    // Add logging middleware
	r.Use(middleware.Recoverer) // Add recovery middleware to handle panics

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthHandler)
	})

	return r
}

func (app *Application) Run(mux http.Handler) error {
	// mux := http.NewServeMux()

	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      mux,
		WriteTimeout: 15 * time.Second, // Set a write timeout of 15 seconds
		ReadTimeout:  15 * time.Second, // Set a read timeout of 15 seconds
		IdleTimeout:  60 * time.Second, // Set an idle timeout of 60 seconds
	}

	log.Printf("Server has started at %s", app.Config.Addr)

	return srv.ListenAndServe()
}
