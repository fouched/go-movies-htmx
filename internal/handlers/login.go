package handlers

import (
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	templates := []string{"/pages/login.gohtml", "/components/alert.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{})
}

func (a *HandlerConfig) LoginPost(w http.ResponseWriter, r *http.Request) {

	pe := r.ParseForm()
	//TODO: we need proper error handling
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
	}

	email := r.Form.Get("email")
	data := make(map[string]interface{})

	if email == "a" {
		a.App.Session.Put(r.Context(), "HasAdmin", true)
		// Good practice: prevent a post re-submit with a http redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "Invalid credentials",
		}
	}

	templates := []string{"/pages/login.gohtml", "/components/alert.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}
