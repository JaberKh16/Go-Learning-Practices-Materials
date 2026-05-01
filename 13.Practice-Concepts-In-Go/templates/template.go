package templates

import (
	"net/http"
	"text/template"
)

type TemplateData struct {
	Title string
	Content string
}


func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	templateData := TemplateData{
		Title:   "Hello World",
		Content: "This is a test",
	}
	
	// TODO: Render the template with the data
	t, err := template.ParseFiles("templates/files/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, templateData)
}
