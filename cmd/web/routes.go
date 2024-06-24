package main

import (
	"github.com/fouched/go-movies-htmx/internal/handlers"
	"github.com/fouched/go-movies-htmx/internal/handlers/admin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Instance.Home)
	mux.Get("/login", handlers.Instance.ShowLogin)
	mux.Post("/login", handlers.Instance.LoginPost)
	mux.Get("/logout", handlers.Instance.Logout)
	mux.Get("/movies", handlers.AllMovies)
	mux.Get("/movies/{id}", handlers.Movie)
	mux.Get("/genres", handlers.Genres)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)
		mux.Get("/movie", admin.EditMovie)
		mux.Get("/catalogue", admin.Catalogue)
		mux.Get("/graphql", admin.GraphQL)
	})

	return mux
}
