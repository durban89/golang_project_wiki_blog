package render

import (
	"html/template"
	"net/http"

	"github.com/durban89/wiki/config"
)

// Render 视图
func Render(w http.ResponseWriter, viewName string, data interface{}) {
	// 视图渲染
	t, err := template.ParseFiles(append(config.CommonTemplatesFiles, config.TemplateDir+"/"+viewName)...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, viewName, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
