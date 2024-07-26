package handlers

import (
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"github.com/fouched/go-movies-htmx/internal/validation"
	"net/http"
	"strconv"
	"time"
)

// AdminMovieAddGet renders the movie page
func (a *HandlerConfig) AdminMovieAddGet(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["title"] = "Add Movie"
	stringMap["action"] = "/admin/movies/add"

	data := make(map[string]interface{})
	data["ratings"] = getRatings("")

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

func getRatings(s string) []models.SelectOption {

	var ratings []models.SelectOption
	ratings = append(ratings, models.SelectOption{Value: "G", Text: "G", Selected: s == "G"})
	ratings = append(ratings, models.SelectOption{Value: "PG", Text: "PG", Selected: s == "PG"})
	ratings = append(ratings, models.SelectOption{Value: "PG13", Text: "PG13", Selected: s == "PG13"})
	ratings = append(ratings, models.SelectOption{Value: "R", Text: "R", Selected: s == "R"})
	ratings = append(ratings, models.SelectOption{Value: "18A", Text: "18A", Selected: s == "18A"})
	return ratings
}

func (a *HandlerConfig) AdminMovieAddPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	runtime, err := strconv.Atoi(r.Form.Get("runtime"))
	if err != nil {
		fmt.Println(err)
	}

	rd := r.Form.Get("releaseDate")
	layout := "2006-01-02"
	releaseDate, err := time.Parse(layout, rd)
	if err != nil {
		fmt.Println(err)
	}

	// read form data
	movie := models.Movie{
		Title:       r.Form.Get("title"),
		Description: r.Form.Get("description"),
		RunTime:     runtime,
		ReleaseDate: releaseDate,
	}

	form := validation.New(r.PostForm)
	form.Required("title", "releaseDate", "runtime", "mpaaRating", "description")

	// deal with validation errors
	if !form.Valid() {
		data := make(map[string]interface{})
		data["ratings"] = getRatings(r.Form.Get("mpaaRating"))
		data["movie"] = movie
		genres, _ := repo.GetAllGenres()
		data["Genres"] = genres

		// re-render the form that did not pass validation
		templates := []string{"/pages/admin/movie.add.gohtml"}
		render.Templates(w, r, templates, true, &models.TemplateData{
			Form: form,
			Data: data,
		})

		return
	}

	//TODO loop through genres

}
