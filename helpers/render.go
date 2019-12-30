package helpers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/durban89/wiki/config"
)

/*
 * @Author: durban
 * @Date:   2019-12-02 18:42:04
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-30 16:38:40
 */

// Render 视图
func Render(w http.ResponseWriter, viewName string, data interface{}) {
	var templateFiles = config.CommonTemplatesFiles
	templateFiles = append(templateFiles, config.TemplateDir+"/"+viewName)

	// 视图渲染
	t, err := template.ParseFiles(templateFiles...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, getLast(viewName), data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func getLast(viewName string) string {
	viewNameArr := strings.Split(viewName, "/")
	return viewNameArr[len(viewNameArr)-1]
}
