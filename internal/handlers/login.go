package handlers

import (
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"log"
	"net/http"
)

func (a *HandlerConfig) ShowLogin(w http.ResponseWriter, r *http.Request) {

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

// LoginPost handles logging the user in
func (a *HandlerConfig) LoginPost(w http.ResponseWriter, r *http.Request) {
	// good practice renew token on login
	_ = a.App.Session.RenewToken(r.Context())

	data := make(map[string]interface{})
	pe := r.ParseForm()

	if pe != nil {
		fmt.Println("Cannot parse form", pe)
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "An unexpected error occurred, please try again later.",
		}

		templates := []string{"/pages/login.gohtml"}
		render.Templates(w, r, templates, true, &models.TemplateData{
			Data: data,
		})
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	id, _, err := repo.Authenticate(email, password)
	if err != nil {
		log.Println(err)
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "Invalid credentials",
		}

		templates := []string{"/pages/login.gohtml"}
		render.Templates(w, r, templates, true, &models.TemplateData{
			Data: data,
		})
		return
	}

	a.App.Session.Put(r.Context(), "userId", id)
	// Good practice: prevent a post re-submit with a http redirect
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

// Logout logs a user out
func (a *HandlerConfig) Logout(w http.ResponseWriter, r *http.Request) {

	_ = a.App.Session.Destroy(r.Context())
	_ = a.App.Session.RenewToken(r.Context())
	templates := []string{"/pages/home.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{})
}
