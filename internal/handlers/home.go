package handlers

import (
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"net/http"
)

func (a *HandlerConfig) Home(w http.ResponseWriter, r *http.Request) {

	templates := []string{"/pages/home.gohtml"}

	render.Templates(w, r, templates, true, &models.TemplateData{})
}
