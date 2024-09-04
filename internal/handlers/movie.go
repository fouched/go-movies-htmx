package handlers

import (
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

var moviesData = make(map[string]interface{})

func (a *HandlerConfig) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := repo.AllMovies()
	if err != nil {
		HandleUnexpectedError(err, w, r)
		return
	}

	data := make(map[string]interface{})
	data["Movies"] = movies
	templates := []string{"/pages/movies.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}

func (a *HandlerConfig) Movie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		HandleUnexpectedError(err, w, r)
		return
	}

	data := make(map[string]interface{})
	movie, err := repo.GetMovieByID(id)
	if err != nil {
		HandleUnexpectedError(err, w, r)
		return
	}

	data["Movie"] = movie
	templates := []string{"/pages/movie.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}

func (a *HandlerConfig) AllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		HandleUnexpectedError(err, w, r)
		return
	}
	movies, err := repo.AllMovies(id)
	if err != nil {
		HandleUnexpectedError(err, w, r)
		return
	}

	data := make(map[string]interface{})
	data["Movies"] = movies
	templates := []string{"/pages/movies.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}
