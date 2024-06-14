package handlers

import (
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"net/http"
)

func (a *HandlerConfig) Home(w http.ResponseWriter, r *http.Request) {

	hasAdmin := a.App.Session.Get(r.Context(), "HasAdmin").(bool)

	boolMap := make(map[string]bool)
	boolMap["HasAdmin"] = hasAdmin

	templates := []string{"/pages/home.gohtml"}

	render.Templates(w, r, templates, true, &models.TemplateData{
		BoolMap: boolMap,
	})
}
