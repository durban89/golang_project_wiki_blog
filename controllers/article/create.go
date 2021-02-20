package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:55:18
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2021-02-20 14:38:37
 */

import (
	"log"
	"net/http"
	"time"

	"wiki/models"
	"wiki/views"
)

// Create 文件
func Create(w http.ResponseWriter, r *http.Request) {
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

	msg := r.URL.Query().Get("err_msg")

	t := time.Now()
	timeStr := t.Format("2006-01-02 03:04:05")

	cate := getArticleCategories()

	log.Println(cate)

	// 视图渲染
	views.Render(w, "article/create.html", struct {
		Msg  string
		Time string
		Cate []models.SelectResult
	}{
		Msg:  msg,
		Time: timeStr,
		Cate: cate,
	})

	return
}
