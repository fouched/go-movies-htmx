package helpers

import (
	"github.com/fouched/go-movies-htmx/internal/config"
	"net/http"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "userId")
	return exists
}
