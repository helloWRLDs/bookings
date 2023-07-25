package main

import (
	"net/http"
	"github.com/supWRLD/bookings/pckg/config"
	"github.com/supWRLD/bookings/pckg/handlers"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
