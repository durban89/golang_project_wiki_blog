package article

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/durban89/wiki/config"
	"github.com/durban89/wiki/helpers"
	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
)

// Save 文章
func Save(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println("post save")
	return
	title := r.FormValue("title")
	body := r.FormValue("body")
	author := r.FormValue("author")
	category := r.FormValue("category")
	ctime := r.FormValue("ctime")
	openStatus := r.FormValue("openStatus")
	token := r.FormValue("token")
	channel := r.Form["channel"]
	method := r.Method

	fmt.Println("author ", author)
	fmt.Println("category ", category)
	fmt.Println("ctime ", ctime)
	fmt.Println("method: ", method)
	fmt.Println("openStatus: ", openStatus)
	fmt.Println("channels: ", channel)
	fmt.Println("token: ", token)

	if len(r.Form.Get("author")) == 0 {
		fmt.Println("author is empty")
	}

	slice := []string{"php", "java", "golang"}
	if !helpers.ValidateInArray(slice, category) {
		fmt.Println("category not in slice")
	}

	chennelSlice := []string{"news", "technology", "other"}
	a := helpers.ValidateSliceIntersection(channel, chennelSlice)
	if len(a) == 0 {
		fmt.Println("channel is empty")
	}

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

	p := &helpers.Page{
		Title: title,
		Body:  []byte(body),
	}

	// 开放状态验证
	openStatusSlice := []string{"1", "2"}

	if !helpers.ValidateInArray(openStatusSlice, openStatus) {
		fmt.Println("openStatus not in openStatusSlice")
	}

	err := p.Save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// Update 更新文章
func Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.NotFound(w, r)
		return
	}

	// model 数据检索查询
	// var articleModel article.Article

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
	t, err := template.ParseFiles(config.TemplateDir + "/update.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, struct {
		Autokid int64
		Title   string
		Content sql.NullString
	}{
		Autokid: autokid,
		Title:   title,
		Content: content,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

// View 文章详情
func View(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.NotFound(w, r)
		return
	}

	// model 数据检索查询
	// var articleModel article.Instance

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
	t, err := template.ParseFiles(config.TemplateDir + "/view.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, struct {
		Autokid int64
		Title   string
		Content string
	}{
		Autokid: autokid,
		Title:   title,
		Content: content.String,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	t, err := template.ParseFiles(config.TemplateDir + "/item.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, struct {
		Data   []models.SelectResult
		Cookie string
	}{
		Data:   qr,
		Cookie: siteName,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
