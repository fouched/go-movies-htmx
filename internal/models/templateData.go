package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	BoolMap         map[string]bool
	Data            map[string]interface{}
	ComponentMap    map[string]interface{}
	CSRFToken       string
	Success         string
	Warning         string
	Error           string
	IsAuthenticated int
}
