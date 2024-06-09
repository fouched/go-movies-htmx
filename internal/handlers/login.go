package handlers

import (
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	templates := []string{"/pages/login.gohtml", "/components/alert.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{})
}

func LoginPost(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	//data["Alert"] = models.Alert{
	//	Class:   "alert-danger",
	//	Message: "Invalid credentials",
	//}

	templates := []string{"/pages/login.gohtml", "/components/alert.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}
