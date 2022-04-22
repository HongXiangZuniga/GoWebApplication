package main

import (
	"net/http"

	"github.com/HongXiangZuniga/GoWebApplication/pkg/config"
	"github.com/HongXiangZuniga/GoWebApplication/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
