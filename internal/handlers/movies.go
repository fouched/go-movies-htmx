package handlers

import (
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

var moviesData = make(map[string]interface{})

func AllMovies(w http.ResponseWriter, r *http.Request) {

	InitInitialState()
	templates := []string{"/pages/movies.gohtml"}

	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: moviesData,
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
		Runtime:     116,
		MppaRating:  "R",
		Description: "Some long description",
	}

	templates := []string{"/pages/movie.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: moviesData,
	})
}

func InitInitialState() {

	layout := "2006-01-02"
	if len(moviesData) == 0 {
		releaseDate, _ := time.Parse(layout, "1986-03-07")
		moviesData[strconv.Itoa(len(moviesData)+1)] = models.Movie{
			ID:          len(moviesData) + 1,
			Title:       "Highlander",
			ReleaseDate: releaseDate,
			Runtime:     116,
			MppaRating:  "R",
			Description: "Some long description",
		}

		releaseDate, _ = time.Parse(layout, "1981-06-12")
		moviesData[strconv.Itoa(len(moviesData)+1)] = models.Movie{
			ID:          len(moviesData) + 1,
			Title:       "Raiders of the lost Ark",
			ReleaseDate: releaseDate,
			Runtime:     115,
			MppaRating:  "PG-13",
			Description: "Some long description",
		}
	}
}
