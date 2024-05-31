package main

import (
	"github.com/fouched/go-movies-htmx/internal/handlers"
	"github.com/fouched/go-movies-htmx/internal/handlers/admin"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/login", handlers.Login)
	mux.Get("/movies", handlers.Movies)
	mux.Get("/genres", handlers.Genres)
	mux.Get("/admin/movie", admin.EditMovie)
	mux.Get("/catalogue", handlers.Catalogue)
	mux.Get("/graphql", handlers.GraphQL)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
