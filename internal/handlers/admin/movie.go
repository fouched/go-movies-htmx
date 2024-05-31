package admin

import (
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"net/http"
)

func EditMovie(w http.ResponseWriter, r *http.Request) {

	templates := []string{"/pages/admin/movie.gohtml"}

	render.Templates(w, r, templates, true, &models.TemplateData{})
}
