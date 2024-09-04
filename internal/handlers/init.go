package handlers

import (
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/config"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"net/http"
)

var Instance *HandlerConfig

type HandlerConfig struct {
	App *config.AppConfig
}

func NewConfig(a *config.AppConfig) *HandlerConfig {
	return &HandlerConfig{
		App: a,
	}
}

func NewHandlers(h *HandlerConfig) {
	Instance = h
}

func HandleUnexpectedError(err error, w http.ResponseWriter, r *http.Request) {

	fmt.Println(err)
	data := make(map[string]interface{})
	data["Alert"] = models.Alert{
		Class:   "alert-danger",
		Message: "An unexpected error occurred, please try again later.",
	}
	templates := []string{"/pages/home.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}
