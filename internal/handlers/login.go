package handlers

import (
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"net/http"
)

func (a *HandlerConfig) Login(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	if a.App.Session.Exists(r.Context(), "AuthError") {
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "Authentication required, please log in",
		}
	}

	templates := []string{"/pages/login.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}

func (a *HandlerConfig) LoginPost(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	pe := r.ParseForm()

	if pe != nil {
		fmt.Println("Cannot parse form", pe)
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "An unexpected error occurred, please try again later.",
		}
	} else {
		email := r.Form.Get("email")

		if email == "a" {
			a.App.Session.Put(r.Context(), "userId", 1)
			// Good practice: prevent a post re-submit with a http redirect
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			data["Alert"] = models.Alert{
				Class:   "alert-danger",
				Message: "Invalid credentials",
			}
		}
	}

	templates := []string{"/pages/login.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})
}

func (a *HandlerConfig) Logout(w http.ResponseWriter, r *http.Request) {

	a.App.Session.Remove(r.Context(), "userId")
	templates := []string{"/pages/home.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{})
}
