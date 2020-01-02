package views

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/durban89/wiki/config"
)

/*
* @Author: durban.zhang
* @Date:   2020-01-02 14:36:57
* @Last Modified by:   durban.zhang
* @Last Modified time: 2020-01-02 14:37:16
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
