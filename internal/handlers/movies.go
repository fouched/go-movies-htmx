package handlers

import (
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

var moviesData = make(map[string]interface{})

func AllMovies(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	movies, err := repo.AllMovies()
	if err != nil {
		fmt.Println(err)
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "An unexpected error occurred, please try again later.",
		}
	} else {
		data["movies"] = movies
	}

	templates := []string{"/pages/movies.gohtml"}

	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}

func Movie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println("Viewing movie with id:", id)

	layout := "2006-01-02"
	releaseDate, _ := time.Parse(layout, "1986-03-07")
	moviesData["movie"] = models.Movie{
		ID:          1,
		Title:       "Another Highlander",
		ReleaseDate: releaseDate,
		RunTime:     116,
		MPAARating:  "R",
		Description: "Some long description",
	}

	templates := []string{"/pages/movie.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: moviesData,
	})
}
