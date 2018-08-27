package helpers

import (
	"html/template"
	"net/http"

	"github.com/durban89/wiki/config"
)

var templates = template.Must(template.ParseFiles(config.TemplateDir+"/edit.html", config.TemplateDir+"/view.html", config.TemplateDir+"/upload.html"))

// RenderTemplate 渲染模板
func RenderTemplate(w http.ResponseWriter, templateName string, p *Page) {
	err := templates.ExecuteTemplate(w, templateName+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
