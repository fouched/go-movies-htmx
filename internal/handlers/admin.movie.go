package handlers

import (
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"github.com/fouched/go-movies-htmx/internal/validation"
	"net/http"
)

// AdminMovieAddGet renders the movie page
func (a *HandlerConfig) AdminMovieAddGet(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["title"] = "Add Movie"
	stringMap["action"] = "/admin/movies/add"

	data := make(map[string]interface{})
	data["ratings"] = getRatings()

	genres, err := repo.GetAllGenres()
	if err != nil {
		fmt.Println(err)
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "An unexpected error occurred, please try again later.",
		}
	} else {
		data["Genres"] = genres
	}

	templates := []string{"/pages/admin/movie.add.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Form:      validation.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
}

func getRatings() []models.SelectOption {

	var ratings []models.SelectOption
	ratings = append(ratings, models.SelectOption{Value: "G", Text: "G"})
	ratings = append(ratings, models.SelectOption{Value: "PG", Text: "PG"})
	ratings = append(ratings, models.SelectOption{Value: "PG13", Text: "PG13"})
	ratings = append(ratings, models.SelectOption{Value: "R", Text: "R"})
	ratings = append(ratings, models.SelectOption{Value: "18A", Text: "18A"})
	return ratings
}

func (a *HandlerConfig) AdminMovieAddPost(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "Add Movie"
	stringMap["action"] = "/admin/movies/add"

	data := make(map[string]interface{})
	data["ratings"] = getRatings()

	genres, err := repo.GetAllGenres()
	if err != nil {
		fmt.Println(err)
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "An unexpected error occurred, please try again later.",
		}
	} else {
		data["Genres"] = genres
	}

	templates := []string{"/pages/admin/movie.add.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Form:      validation.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
}
