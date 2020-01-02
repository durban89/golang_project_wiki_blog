package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:54:35
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 17:08:50
 */

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
	"github.com/durban89/wiki/models/articletag"
	"github.com/durban89/wiki/views"
)

// Update 更新文章
func Update(w http.ResponseWriter, r *http.Request) {
	session, error := SessionManager.SessionStart(w, r)
	if error != nil {
		http.Error(w, "SessionStart Fail", 403)
		return
	}

	userID := session.Get("user_id")

	if userID == nil {
		http.Redirect(w, r, "/auth/login", http.StatusFound)
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		http.NotFound(w, r)
		return
	}

	var articleID int64
	var title string
	var content sql.NullString
	var summary sql.NullString
	var categoryID int64
	var createdAt string

	// selectField := models.SelectValues{
	// 	"autokid":     &articleID,
	// 	"title":       &title,
	// 	"content":     &content,
	// 	"summary":     &summary,
	// 	"category_id": &categoryID,
	// 	"created_at":  &createdAt,
	// }
	selectField := []string{
		"autokid",
		"title",
		"content",
		"summary",
		"category_id",
		"created_at",
	}

	where := models.WhereValues{
		"autokid": models.WhereCondition{
			Operator: "=",
			Value:    id,
		},
	}

	_, err := article.Instance.QueryOne(selectField, where)

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
			Value:    id,
		},
	}

	order := models.OrderValues{
		"autokid": models.OrderCondition{
			OrderBy: "DESC",
		},
	}

	tags, err := articletag.Instance.Query(selectTagField, whereTag, order, 0, 100)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	var tagsArr = []string{}

	for _, v := range tags {
		tagsArr = append(tagsArr, v["name"].(string))
	}

	cate := getArticleCategories()

	log.Println(cate)

	// 视图渲染
	views.Render(w, "article/update.html", struct {
		Autokid    int64
		Title      string
		Content    string
		Summary    string
		CategoryID int64
		CreatedAt  string
		Tags       string
		Cate       []models.SelectResult
	}{
		Autokid:    articleID,
		Title:      title,
		Content:    content.String,
		Summary:    summary.String,
		CategoryID: categoryID,
		CreatedAt:  createdAt,
		Tags:       strings.Join(tagsArr, ";"),
		Cate:       cate,
	})

	return
}
