package helpers

import (
	"html/template"
	"net/http"

	"github.com/durban89/wiki/config"
	"github.com/durban89/wiki/models"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles(
		config.TemplateDir + "/index.html"))
}

// RenderTemplate 渲染模板
func RenderTemplate(w http.ResponseWriter, templateName string, p *Page) {
	err := templates.ExecuteTemplate(w, templateName+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RenderTemplateWithItemData 多数据渲染模板
func RenderTemplateWithItemData(w http.ResponseWriter, templateName string, data []models.SelectResult) {
	err := templates.ExecuteTemplate(w, templateName+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
