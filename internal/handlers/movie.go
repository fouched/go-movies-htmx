package handlers

import (
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

var moviesData = make(map[string]interface{})

func (a *HandlerConfig) AllMovies(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	movies, err := repo.AllMovies()
	if err != nil {
		fmt.Println(err)
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "An unexpected error occurred, please try again later.",
		}
	} else {
		data["Movies"] = movies
	}

	templates := []string{"/pages/movies.gohtml"}

	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}

func (a *HandlerConfig) Movie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	movieID, err := strconv.Atoi(id)

	data := make(map[string]interface{})

	movie, err := repo.GetMovieByID(movieID)
	if err != nil {
		fmt.Println(err)
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "An unexpected error occurred, please try again later.",
		}
	} else {
		data["Movie"] = movie
	}

	templates := []string{"/pages/movie.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}
