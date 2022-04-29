package server

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseGlob("views/*"))

func RenderTemplate(w http.ResponseWriter, tmpl string, data Page) {
	err := templ.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
