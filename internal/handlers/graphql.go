package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/graph"
	"github.com/fouched/go-movies-htmx/internal/models"
	"github.com/fouched/go-movies-htmx/internal/render"
	"github.com/fouched/go-movies-htmx/internal/repo"
	"net/http"
)

type ResultList struct {
	List   []models.Movie
	Search []models.Movie
}

type Result struct {
	Data ResultList
}

func (a *HandlerConfig) GraphQLGet(w http.ResponseWriter, r *http.Request) {

	templates := []string{"/pages/graphql.gohtml"}
	render.Templates(w, r, templates, true, &models.TemplateData{})
}

func (a *HandlerConfig) GraphQLPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	payload := `
        {
            list {
                id
                title
                runtime
                releasedate
                mpaarating
            }
        }`

	search := r.Form.Get("search")
	if search != "" {
		payload = fmt.Sprintf(`
        {
            search(titleContains: "%s") {
                id
                title
                runtime
                releasedate
                mpaarating
            }
        }`, search)
	}

	// populate Graph type with the movies
	movies, _ := repo.AllMovies() // this is horrible, cache it...
	g := graph.New(movies)
	g.QueryString = payload

	// perform the query
	resp, err := g.Query()
	if err != nil {
		HandleUnexpectedError(err, w, r)
		return
	}
	j, _ := json.MarshalIndent(resp, "", "\t")
	//fmt.Println(string(j))

	var result Result
	err = json.Unmarshal(j, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	if result.Data.List != nil {
		data["Movies"] = result.Data.List
	} else {
		data["Movies"] = result.Data.Search
	}

	templates := []string{"/snippets/graphResults.gohtml"}
	render.Templates(w, r, templates, false, &models.TemplateData{
		Data: data,
	})
}
