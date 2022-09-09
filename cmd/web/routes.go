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

	//MiddleWare
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
