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

var data = make(map[string]interface{})

func AllMovies(w http.ResponseWriter, r *http.Request) {

	InitInitialState()
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
	data["movie"] = models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: releaseDate,
		Runtime:     116,
		MppaRating:  "R",
		Description: "Some long description",
	}

	templates := []string{"/pages/movie.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}

func InitInitialState() {

	layout := "2006-01-02"
	if len(data) == 0 {
		releaseDate, _ := time.Parse(layout, "1986-03-07")
		data[strconv.Itoa(len(data)+1)] = models.Movie{
			ID:          len(data) + 1,
			Title:       "Highlander",
			ReleaseDate: releaseDate,
			Runtime:     116,
			MppaRating:  "R",
			Description: "Some long description",
		}

		releaseDate, _ = time.Parse(layout, "1981-06-12")
		data[strconv.Itoa(len(data)+1)] = models.Movie{
			ID:          len(data) + 1,
			Title:       "Raiders of the lost Ark",
			ReleaseDate: releaseDate,
			Runtime:     115,
			MppaRating:  "PG-13",
			Description: "Some long description",
		}
	}
}
