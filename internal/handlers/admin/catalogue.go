package admin

import (
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"net/http"
)

func Catalogue(w http.ResponseWriter, r *http.Request) {

	templates := []string{"/pages/admin/catalogue.gohtml"}

	render.Templates(w, r, templates, true, &models.TemplateData{})
}