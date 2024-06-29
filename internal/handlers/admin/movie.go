package admin

import (
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/validation"
	"net/http"
)

// Movie renders the movie page
func Movie(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["title"] = "Add Movie"
	stringMap["action"] = "/admin/movie/add"

	data := make(map[string]interface{})
	data["ratings"] = getRatings()

	templates := []string{"/pages/admin/movie.gohtml"}
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
	ratings = append(ratings, models.SelectOption{Value: "18", Text: "18"})
	return ratings
}
