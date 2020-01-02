package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:53:27
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-31 19:00:31
 */

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/durban89/wiki/helpers"
	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
)

// View 文章详情
func View(w http.ResponseWriter, r *http.Request) {

	session, error := SessionManager.SessionStart(w, r)

	if error != nil {
		log.Println(error)
		http.Error(w, "Session启动失败", 500)
		return
	}

	userID := session.Get("user_id")

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
	var authorID string

	articleID = 0

	selectField := models.SelectValues{
		"autokid":     &articleID,
		"title":       &title,
		"category_id": &categoryID,
		"author_id":   &authorID,
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

	tagsArr := getArticleTag(w, r, id)
	author, err := getAuthor(authorID)
	category := getArticleCategory(strconv.FormatInt(categoryID, 10))

	if err != nil {
		http.Error(w, "用户信息查询失败", 500)
		return
	}

	// 视图渲染
	helpers.Render(w, "article/view.html", struct {
		Autokid    string
		Title      string
		Content    template.HTML
		CategoryID int64
		Category   map[string]string
		CreatedAt  string
		Tags       []string
		UserID     interface{}
		Author     map[string]string
	}{
		Autokid:    id,
		Title:      title,
		Content:    template.HTML(content.String),
		CategoryID: categoryID,
		Category:   category,
		CreatedAt:  createdAt,
		Tags:       tagsArr,
		UserID:     userID,
		Author:     author,
	})

	return
}
