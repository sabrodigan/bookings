package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sabrodigan/bookings/pkg/config"
	"github.com/sabrodigan/bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Index)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/add", handlers.Repo.Add)
	mux.Get("/create", handlers.Repo.Create)
	mux.Get("/delete", handlers.Repo.Delete)
	mux.Get("/divide", handlers.Repo.Divide)
	mux.Get("/move", handlers.Repo.Move)
	mux.Get("/search", handlers.Repo.Search)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
