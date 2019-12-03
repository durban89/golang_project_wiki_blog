package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:54:35
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-02 18:46:55
 */

import (
	"database/sql"
	"net/http"

	"github.com/durban89/wiki/helpers"
	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
)

// Update 更新文章
func Update(w http.ResponseWriter, r *http.Request) {
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
		http.NotFound(w, r)
		return
	}

	// 视图渲染
	helpers.Render(w, "update.html", struct {
		Autokid int64
		Title   string
		Content string
	}{
		Autokid: autokid,
		Title:   title,
		Content: content.String,
	})

	return
}
