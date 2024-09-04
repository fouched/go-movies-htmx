package handlers

import (
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"net/http"
)

func (a *HandlerConfig) Genres(w http.ResponseWriter, r *http.Request) {
	genres, err := repo.AllGenres()
	if err != nil {
		HandleUnexpectedError(err, w, r)
		return
	}

	data := make(map[string]interface{})
	data["Genres"] = genres
	templates := []string{"/pages/genres.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}
