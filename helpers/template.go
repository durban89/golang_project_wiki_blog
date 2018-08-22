package helpers

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html", "templates/upload.html"))

// RenderTemplate 渲染模板
func RenderTemplate(w http.ResponseWriter, templateName string, p *Page) {
	err := templates.ExecuteTemplate(w, templateName+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
