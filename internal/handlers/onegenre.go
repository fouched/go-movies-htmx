package handlers

import (
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (a *HandlerConfig) OneGenre(w http.ResponseWriter, r *http.Request) {

	stringmap := make(map[string]string)
	stringmap["GenreName"] = chi.URLParam(r, "genreName")

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
	templates := []string{"/pages/onegenre.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data:      data,
		StringMap: stringmap,
	})
}
