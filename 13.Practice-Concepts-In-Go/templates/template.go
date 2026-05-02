// use case of html/template package
package templates

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
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
	// func(t *Template) Execute(wr io.Writer, data interface{}) error
	t.Execute(w, templateData)
}

func EscappedOrUnscappedHTMLTags(w http.ResponseWriter, r *http.Request) {
	templateData := TemplateData{
		Title:   "Hello World",
		Content: "This is a test <script>alert('XSS')</script>",
	}
	
	// escapped the srting
	fmt.Println(template.HTMLEscapeString(templateData.Content))
	// unscapped the srting
	fmt.Println(template.HTML(templateData.Content))
	
}

func renderConditionalTemplating() {
	tmplString := "{{ if .Title }}{{ .Title }}{{ else }}Default Title{{ end }}" // basically checking if the title is empty or not
	// template.Must() will panic if there is an error
	tmpl := template.Must(template.New("conditional").Parse(tmplString))	
	//func (t *Template) Execute(wr io.Writer, data interface{}) error
	err := tmpl.Execute(os.Stdout, 1)
	if err != nil {
		fmt.Println(err)
	}

	// render product list
	productList := []string{"Product 1", "Product 2", "Product 3"}
	tmplString = "{{ range . }}{{ . }}{{ end }}"
	tmpl = template.Must(template.New("productList").Parse(tmplString))
	err = tmpl.Execute(os.Stdout, productList)
	if err != nil {
		fmt.Println(err)
	}
}
