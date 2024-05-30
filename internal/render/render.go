package render

import (
	"fmt"
	"github.com/fouched/go-movies-htmx/internal/models"
	"html/template"
	"net/http"
)

var pathToTemplates = "./templates"

// Templates can render multiple templates. "Parent" templates should be defined first
func Templates(w http.ResponseWriter, r *http.Request, tmpl []string, addBaseTemplate bool, td *models.TemplateData) {

	for i, t := range tmpl {
		tmpl[i] = pathToTemplates + t
	}

	if addBaseTemplate {
		tmpl = append(tmpl, pathToTemplates+"/base.layout.gohtml")
	}

	parsedTemplate, _ := template.ParseFiles(tmpl...)
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}

}
