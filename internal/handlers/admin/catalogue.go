package admin

import (
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"net/http"
)

func Catalogue(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	movies, err := repo.AllMovies()
	if err != nil {
		fmt.Println(err)
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "An unexpected error occurred, please try again later.",
		}
	} else {
		data["movies"] = movies
	}

	templates := []string{"/pages/admin/catalogue.gohtml"}

	render.Templates(w, r, templates, true, &models.TemplateData{
		Data: data,
	})

}
