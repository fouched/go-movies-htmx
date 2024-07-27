package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"github.com/fouched/go-movies-htmx/internal/validation"
	"io"
	"log"
	"net/http"
	"net/url"
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

	genres, _ := repo.GetAllGenres()
	selectedGenres := r.Form["genres"]

	// read form data
	movie := models.Movie{
		Title:       r.Form.Get("title"),
		Description: r.Form.Get("description"),
		RunTime:     runtime,
		ReleaseDate: releaseDate,
		Genres:      genres,
	}

	form := validation.New(r.PostForm)

	// deal with validation errors
	form.Required("title", "releaseDate", "runtime", "mpaaRating", "description")

	if len(selectedGenres) == 0 {
		form.Errors.Add("genres", "Please select a genre")
	} else {
		for _, genre := range genres {
			for _, selectedGenre := range selectedGenres {
				sg, _ := strconv.Atoi(selectedGenre)
				if genre.ID == sg {
					genre.Checked = true
				}
			}
		}
	}

	if !form.Valid() {
		data := make(map[string]interface{})
		data["ratings"] = getRatings(r.Form.Get("mpaaRating"))
		data["movie"] = movie
		data["Genres"] = genres

		// re-render the form that did not pass validation
		templates := []string{"/pages/admin/movie.add.gohtml"}
		render.Templates(w, r, templates, true, &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// validation passed persist the form

	// try to get an image
	movie = a.getPoster(movie)

	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()
	movie.MPAARating = r.Form.Get("mpaaRating")

	// insert the movie
	newID, err := repo.InsertMovie(&movie)
	if err != nil {
		fmt.Println(err)
	}

	// handle genres
	var ga []int
	for _, selectedGenre := range selectedGenres {
		sg, _ := strconv.Atoi(selectedGenre)
		ga = append(ga, sg)
	}
	movie.GenresArray = ga

	err = repo.UpdateMovieGenres(newID, movie.GenresArray)
	if err != nil {
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
		return
	}

	// Good practice: prevent a post re-submit with a http redirect
	http.Redirect(w, r, "/admin/catalogue", http.StatusSeeOther)

}

func (a *HandlerConfig) getPoster(movie models.Movie) models.Movie {
	type TheMovieDB struct {
		Page    int `json:"page"`
		Results []struct {
			PosterPath string `json:"poster_path"`
		} `json:"results"`
	}

	client := http.Client{}
	theUrl := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s", a.App.APIKey)

	req, err := http.NewRequest("GET", theUrl+"&query="+url.QueryEscape(movie.Title), nil)
	if err != nil {
		// api does not have the movie
		log.Println(err)
		return movie
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return movie
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return movie
	}

	var responseObject TheMovieDB
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		log.Println(err)
		return movie
	}

	if len(responseObject.Results) > 0 {
		movie.Image = responseObject.Results[0].PosterPath
	}

	return movie
}
