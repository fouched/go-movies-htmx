package main

import (
	"github.com/fouched/go-movies-htmx/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/movies", handlers.Movies)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
