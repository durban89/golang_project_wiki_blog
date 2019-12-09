package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:54:35
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-09 20:56:48
 */

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/durban89/wiki/helpers"
	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
	"github.com/durban89/wiki/models/articletag"
)

// Update 更新文章
func Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.NotFound(w, r)
		return
	}

	var articleID int64
	var title string
	var content sql.NullString
	var categoryID int64
	var createdAt string

	selectField := models.SelectValues{
		"autokid":     &articleID,
		"title":       &title,
		"content":     &content,
		"category_id": &categoryID,
		"created_at":  &createdAt,
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

	var tagName string

	selectTagField := models.SelectValues{
		"name": &tagName,
	}

	whereTag := models.WhereValues{
		"article_id": models.WhereCondition{
			Operator: "=",
			Value:    strconv.FormatInt(articleID, 10),
		},
	}

	tags, err := articletag.Instance.Query(selectTagField, whereTag, 1, 100)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	var tagsArr = []string{}

	for _, v := range tags {
		tagsArr = append(tagsArr, v["name"].(string))
	}

	// 视图渲染
	helpers.Render(w, "update.html", struct {
		Autokid    int64
		Title      string
		Content    string
		CategoryID int64
		CreatedAt  string
		Tags       string
	}{
		Autokid:    articleID,
		Title:      title,
		Content:    content.String,
		CategoryID: categoryID,
		CreatedAt:  createdAt,
		Tags:       strings.Join(tagsArr, ";"),
	})

	return
}
