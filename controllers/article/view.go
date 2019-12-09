package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:53:27
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-02 18:47:05
 */

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/durban89/wiki/helpers"
	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
)

// View 文章详情
func View(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.NotFound(w, r)
		return
	}

	var autokid int64
	var title string
	var content sql.NullString
	selectField := models.SelectValues{
		"autokid": &autokid,
		"title":   &title,
		"content": &content,
	}

	where := models.WhereValues{
		"autokid": models.WhereCondition{
			Operator: "=",
			Value:    id,
		},
	}

	err := article.Instance.QueryOne(selectField, where)

	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}

	// 视图渲染
	helpers.Render(w, "view.html", struct {
		Autokid int64
		Title   string
		Content template.HTML
	}{
		Autokid: autokid,
		Title:   title,
		Content: template.HTML(content.String),
	})

	return
}
