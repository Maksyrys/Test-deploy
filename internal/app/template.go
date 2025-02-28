package app

import "html/template"

type Template struct {
	templates *template.Template
}

func NewTemplate(name string) *Template {
	return &Template{
		templates: template.Must(template.ParseFiles(name)),
	}
}

var templates = NewTemplate("../../templates/index.html")
