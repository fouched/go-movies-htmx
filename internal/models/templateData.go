package models

import "github.com/fouched/go-movies-htmx/internal/validation"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	BoolMap         map[string]bool
	Data            map[string]interface{}
	CSRFToken       string
	Success         string
	Warning         string
	Error           string
	IsAuthenticated int
	Form            *validation.Form
}
