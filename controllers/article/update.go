package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:54:35
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 18:25:09
 */

import (
	"log"
	"net/http"
	"strings"

	"wiki/models"
	"wiki/models/article"
	ArticleTag "wiki/models/article/tag"
	"wiki/views"
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

	where := models.WhereValues{
		"autokid": models.WhereCondition{
			Operator: "=",
			Value:    id,
		},
	}

	articleModel, err := article.Instance.QueryOne([]string{
		"autokid",
		"title",
		"content",
		"summary",
		"category_id",
		"author_id",
		"created_at",
	}, where)

	if err != nil {
		http.NotFound(w, r)
		return
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

	tags, err := ArticleTag.Instance.Query([]string{
		"name",
	}, whereTag, order, 0, 100)

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
		Article models.SelectResult
		Tags    string
		Cate    []models.SelectResult
	}{
		Article: articleModel,
		Tags:    strings.Join(tagsArr, ";"),
		Cate:    cate,
	})

	return
}
