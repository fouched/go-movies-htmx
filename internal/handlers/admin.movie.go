package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"github.com/fouched/go-movies-htmx/internal/validation"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// AdminMovieAddGet renders the add movie page
func (a *HandlerConfig) AdminMovieAddGet(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["Title"] = "Add Movie"
	stringMap["Action"] = "/admin/movies/add"

	data := make(map[string]interface{})

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

	data["Ratings"] = getRatings("")

	templates := []string{"/pages/admin/movie.add.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Form:      validation.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
}

// AdminMovieEditGet renders the add movie page
func (a *HandlerConfig) AdminMovieEditGet(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["Title"] = "Edit Movie"
	stringMap["Action"] = "/admin/movies/edit"

	id := chi.URLParam(r, "id")
	movieID, err := strconv.Atoi(id)

	data := make(map[string]interface{})

	movie, err := repo.GetMovieByID(movieID)
	if err != nil {
		fmt.Println(err)
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "An unexpected error occurred, please try again later.",
		}
	} else {
		data["Movie"] = movie
	}

	allGenres, err := repo.GetAllGenres()

	// pre-select the movie's genres
	for _, genre := range allGenres {
		for _, movieGenre := range movie.Genres {
			if genre.ID == movieGenre.ID {
				genre.Checked = true
			}
		}
	}

	if err != nil {
		fmt.Println(err)
		data["Alert"] = models.Alert{
			Class:   "alert-danger",
			Message: "An unexpected error occurred, please try again later.",
		}
	} else {
		data["Genres"] = allGenres
	}

	data["Ratings"] = getRatings(movie.MPAARating)

	templates := []string{"/pages/admin/movie.add.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{
		Form:      validation.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
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
		MPAARating:  r.Form.Get("mpaaRating"),
	}

	// try to get an image
	movie = a.getPoster(movie)

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
		data["Ratings"] = getRatings(r.Form.Get("mpaaRating"))
		data["Movie"] = movie
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
	// add date fields
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

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

func (a *HandlerConfig) AdminMovieEditPost(w http.ResponseWriter, r *http.Request) {
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

	movieId, _ := strconv.Atoi(r.Form.Get("movieId"))

	// read form data
	movie := models.Movie{
		ID:          movieId,
		Title:       r.Form.Get("title"),
		Description: r.Form.Get("description"),
		RunTime:     runtime,
		ReleaseDate: releaseDate,
		Genres:      genres,
		MPAARating:  r.Form.Get("mpaaRating"),
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
		data["Ratings"] = getRatings(r.Form.Get("mpaaRating"))
		data["Movie"] = movie
		data["Genres"] = genres

		// re-render the form that did not pass validation
		templates := []string{"/pages/admin/movie.add.gohtml"}
		render.Templates(w, r, templates, true, &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// validation passed update the movie
	movie.UpdatedAt = time.Now()
	err = repo.UpdateMovie(&movie)
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

	err = repo.UpdateMovieGenres(movie.ID, movie.GenresArray)
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

func getRatings(s string) []models.SelectOption {
	var ratings []models.SelectOption
	ratings = append(ratings, models.SelectOption{Value: "G", Text: "G", Selected: s == "G"})
	ratings = append(ratings, models.SelectOption{Value: "PG", Text: "PG", Selected: s == "PG"})
	ratings = append(ratings, models.SelectOption{Value: "PG13", Text: "PG13", Selected: s == "PG13"})
	ratings = append(ratings, models.SelectOption{Value: "R", Text: "R", Selected: s == "R"})
	ratings = append(ratings, models.SelectOption{Value: "18", Text: "18", Selected: s == "18"})
	return ratings
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
