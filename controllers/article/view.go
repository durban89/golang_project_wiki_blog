package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:53:27
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-12 16:42:25
 */

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/durban89/wiki/helpers"
	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
	"github.com/durban89/wiki/models/articletag"
)

// View 文章详情
func View(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.NotFound(w, r)
		return
	}

	var articleID int64
	var title string
	var content sql.NullString
	var createdAt string
	var categoryID int64

	articleID = 0

	selectField := models.SelectValues{
		"autokid":     &articleID,
		"title":       &title,
		"category_id": &categoryID,
		"content":     &content,
		"created_at":  &createdAt,
	}

	where := models.WhereValues{
		"autokid": models.WhereCondition{
			Operator: "=",
			Value:    id,
		},
	}

	err := article.Instance.QueryOne(selectField, where)

	if err != nil || articleID == 0 {
		http.NotFound(w, r)
		return
	}

	// tag query
	var tagName string

	selectTagField := models.SelectValues{
		"name": &tagName,
	}

	whereTag := models.WhereValues{
		"article_id": models.WhereCondition{
			Operator: "=",
			Value:    id,
		},
	}

	tags, err := articletag.Instance.Query(selectTagField, whereTag, 0, 100)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	var tagsArr = []string{}

	for _, v := range tags {
		tagsArr = append(tagsArr, v["name"].(string))
	}

	// 视图渲染
	helpers.Render(w, "view.html", struct {
		Autokid    string
		Title      string
		Content    template.HTML
		CategoryID int64
		CreatedAt  string
		Tags       []string
	}{
		Autokid:    id,
		Title:      title,
		Content:    template.HTML(content.String),
		CategoryID: categoryID,
		CreatedAt:  createdAt,
		Tags:       tagsArr,
	})

	return
}
