package helpers

import (
	"html/template"
	"net/http"

	"github.com/durban89/wiki/config"
)

/*
 * @Author: durban
 * @Date:   2019-12-02 18:42:04
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-02 18:42:34
 */

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
