package article

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/durban89/wiki/helpers/render"
	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
)

// Error 错误显示
func Error(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	return
}

// Delete 删除操作
func Delete(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	return
}

// Create 文件
func Create(w http.ResponseWriter, r *http.Request) {
	// 视图渲染
	render.Render(w, "create.html", nil)

	return
}

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
	render.Render(w, "update.html", struct {
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
	render.Render(w, "view.html", struct {
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

// Item 列表
func Item(w http.ResponseWriter, r *http.Request) {
	var siteName string
	cookie, err := r.Cookie("site_name_cookie")

	if err != nil {
		expired := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{
			Name:    "site_name_cookie",
			Value:   "gowhich_cookie",
			Expires: expired,
		}

		http.SetCookie(w, &cookie)
	} else {
		siteName = cookie.Value
	}

	// var articleModel article.Article

	var autokid int64
	var title string
	selectField := models.SelectValues{
		"autokid": &autokid,
		"title":   &title,
	}

	where := models.WhereValues{}

	qr, err := article.Instance.Query(selectField, where, 0, 10)

	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}

	render.Render(w, "item.html", struct {
		Data   []models.SelectResult
		Cookie string
	}{
		Data:   qr,
		Cookie: siteName,
	})

	return
}
