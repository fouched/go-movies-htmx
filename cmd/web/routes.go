package main

import (
	"github.com/fouched/go-movies-htmx/internal/handlers"
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
	mux.Get("/movies", handlers.Instance.AllMovies)
	mux.Get("/movies/{id}", handlers.Instance.Movie)
	mux.Get("/genres", handlers.Instance.Genres)
	mux.Get("/genre/{id}/{genreName}", handlers.Instance.OneGenre)
	mux.Get("/graphql", handlers.Instance.GraphQLGet)
	mux.Post("/graphql", handlers.Instance.GraphQLPost)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)
		mux.Get("/movies/add", handlers.Instance.AdminMovieAddGet)
		mux.Post("/movies/add", handlers.Instance.AdminMovieAddPost)
		mux.Get("/movies/{id}", handlers.Instance.AdminMovieEditGet)
		mux.Post("/movies/edit", handlers.Instance.AdminMovieEditPost)
		mux.Post("/movies/delete", handlers.Instance.AdminMovieDeletePost)
		mux.Get("/catalogue", handlers.Instance.AdminCatalogue)
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
