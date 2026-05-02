// use case of text/template package
package templates

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material  string
	Price float64
	Count int
}

func RenderTextTemplate() {
	inventory := Inventory{
		Material:  "Apple",
		Price: 1.99,
		Count: 10,
	}

	tmpl, err := template.New("test").Parse("We have {{.Count}} {{.Material}} in stock.")
	if err != nil {
		panic(err)
	}
	
	err = tmpl.Execute(os.Stdout, inventory)
	if err != nil {
		panic(err)
	}
}

func RenderTextTemplateWithFunc() {
	inventory := Inventory{
		Material:  "Apple",
		Price: 1.99,
		Count: 10,
	}

	tmpl, err := template.New("test").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}).Parse("We have {{.Count}} {{.Material}} in stock.")
	if err != nil {
		panic(err)
	}
	
	err = tmpl.Execute(os.Stdout, inventory)
	if err != nil {
		panic(err)
	}
}