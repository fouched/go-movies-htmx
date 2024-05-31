package handlers

import (
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"net/http"
)

func GraphQL(w http.ResponseWriter, r *http.Request) {

	templates := []string{"/pages/graphql.gohtml"}

	render.Templates(w, r, templates, true, &models.TemplateData{})
}
