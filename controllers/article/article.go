package article

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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

// Save 文章
func Save(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var id string
	// fmt.Println(r.Form)
	// return
	id = r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	category := r.FormValue("category")

	fmt.Println("category ", category)

	if title == "" || content == "" || category == "" {
		fmt.Println("tishi 错误信息")
	}

	if id != "" {
		update := models.UpdateValues{
			"title":   title,
			"content": content,
		}
		where := models.WhereValues{
			"autokid": models.WhereCondition{
				Operator: "=",
				Value:    id,
			},
		}

		_, err := article.Instance.Update(update, where)

		if err != nil {
			http.Redirect(w, r, "/articles/error", http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Println("crete")
		insert := models.InsertValues{
			"title":   title,
			"content": content,
		}

		insertID, err := article.Instance.Create(insert)

		fmt.Println(insertID)

		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/articles/error", http.StatusSeeOther)
			return
		}

		id = strconv.FormatInt(insertID, 10)
	}

	if id == "" {
		http.Redirect(w, r, "/articles/view/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/articles/view/?id="+id, http.StatusSeeOther)
	}

	// return
	// if len(r.Form.Get("author")) == 0 {
	// 	fmt.Println("author is empty")
	// }

	// slice := []string{"php", "java", "golang"}
	// if !helpers.ValidateInArray(slice, category) {
	// 	fmt.Println("category not in slice")
	// }

	// chennelSlice := []string{"news", "technology", "other"}
	// a := helpers.ValidateSliceIntersection(channel, chennelSlice)
	// if len(a) == 0 {
	// 	fmt.Println("channel is empty")
	// }

	// XSS Example
	// fmt.Println("author: ", template.HTMLEscapeString(r.Form.Get("author"))) // print at server side
	// fmt.Println("author: ", template.HTMLEscapeString(r.Form.Get("author")))
	// template.HTMLEscape(w, []byte(r.Form.Get("author"))) // responded to clients
	// return

	// 重复提交 Example
	// if token != "" {
	// 	// check token validity
	// 	fmt.Println("To Validate Token")
	// } else {
	// 	// give error if no token
	// 	fmt.Println("Token is empty")
	// }

	// p := &helpers.Page{
	// 	Title: title,
	// 	Body:  []byte(body),
	// }

	// // 开放状态验证
	// openStatusSlice := []string{"1", "2"}

	// if !helpers.ValidateInArray(openStatusSlice, openStatus) {
	// 	fmt.Println("openStatus not in openStatusSlice")
	// }

	// err := p.Save()

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// http.Redirect(w, r, "/articles/view/?id="+id, http.StatusFound)
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
