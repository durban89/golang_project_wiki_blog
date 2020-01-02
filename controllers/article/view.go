package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:53:27
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 16:16:03
 */

import (
	"log"
	"net/http"
	"strconv"

	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
	"github.com/durban89/wiki/views"
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

	where := models.WhereValues{
		"autokid": models.WhereCondition{
			Operator: "=",
			Value:    id,
		},
	}

	articleModel, err := article.Instance.QueryOne(nil, where)

	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	tagsArr := getArticleTag(w, r, id)
	author, err := getAuthor(strconv.FormatInt(articleModel["author_id"].(int64), 10))

	if err != nil {
		log.Println(err)
		http.Error(w, "用户信息查询失败", 500)
		return
	}

	category := getArticleCategory(strconv.FormatInt(articleModel["category_id"].(int64), 10))

	// 视图渲染
	views.Render(w, "article/view.html", struct {
		Autokid  string
		Article  models.SelectResult
		Category models.SelectResult
		Tags     []string
		UserID   interface{}
		Author   models.SelectResult
	}{
		Autokid:  id,
		Article:  articleModel,
		Category: category,
		Tags:     tagsArr,
		UserID:   userID,
		Author:   author,
	})

	return
}
