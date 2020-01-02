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
 * @Last Modified time: 2020-01-02 18:16:12
 */

// Render 视图
func Render(w http.ResponseWriter, viewName string, data interface{}) {
	templateName := getLast(viewName)

	// map func
	funcMap := template.FuncMap{
		"ToString": ToString,
		"ToInt":    ToInt,
		"ToInt64":  ToInt64,
	}
	t := template.New(templateName).Funcs(funcMap)

	// template parse
	var templateFiles = config.CommonTemplatesFiles
	templateFiles = append(templateFiles, config.TemplateDir+"/"+viewName)
	t = template.Must(t.ParseFiles(templateFiles...))

	// 视图渲染
	// t, err := template.ParseFiles(templateFiles...)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// t.Funcs(template.FuncMap{"tostring": ToString})

	err := t.ExecuteTemplate(w, templateName, data)

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
