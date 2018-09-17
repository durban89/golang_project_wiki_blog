package controllers

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/durban89/wiki/config"

	"github.com/durban89/wiki/helpers"
	"github.com/durban89/wiki/models"
)

// ArticleItem 列表
func ArticleItem(w http.ResponseWriter, r *http.Request) {
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

	var blogModel models.BlogModel

	var autokid int64
	var title string
	selectField := models.SelectValues{
		"autokid": &autokid,
		"title":   &title,
	}

	where := models.WhereValues{}

	qr, err := blogModel.Query(selectField, where, 0, 10)

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

// ArticleViewWithID 获取文章的id
func ArticleViewWithID(w http.ResponseWriter, r *http.Request) {
	if strings.ToLower(r.Method) == "get" {
		var validPath = regexp.MustCompile("^/(view|edit)/([a-zA-Z0-9]+)$")
		m := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Println(m)
		if m == nil {
			http.NotFound(w, r)
			return
		}

		// 获取文章标题或者文章ID
		fmt.Println(m[2:])

		fmt.Fprintf(w, "Welcome to the home page!")
		return
	}

	http.NotFound(w, r)
	return

}

// ArticleView 查看文章
func ArticleView(w http.ResponseWriter, r *http.Request, title string) {
	p, err := helpers.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	// p.Script = "<script>alert('you have been pwned')</script>"
	// p.Html = template.HTML("<script>alert('you have been pwned')</script>")
	helpers.RenderTemplate(w, "view", p)
}

// ArticleSave 存储文章
func ArticleSave(w http.ResponseWriter, r *http.Request, title string) {
	r.ParseForm()

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

// ArticleEdit 编辑文章
func ArticleEdit(w http.ResponseWriter, r *http.Request) {
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

	update := map[string]string{}

	var blogModel models.BlogModel

	if strings.ToLower(r.Method) == "get" {
		var autokid int64
		var title string
		selectField := models.SelectValues{
			"autokid": &autokid,
			"title":   &title,
		}

		err := blogModel.QueryOne(selectField, where)

		if err != nil {
			http.NotFound(w, r)
			return
		}

		p := helpers.Page{
			Title: title,
			ID:    autokid,
		}

		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		p.Token = token
		helpers.RenderTemplate(w, "edit", &p)
	} else if strings.ToLower(r.Method) == "post" {
		title := r.FormValue("title")
		if title == "" {
			http.Redirect(w, r, fmt.Sprintf("/edit?id=%s", id), http.StatusFound)
			return
		}

		update["title"] = title

		var blogInstance models.BlogModel
		_, err := blogInstance.Update(update, where)

		if err != nil {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/edit?id=%s", id), http.StatusFound)
		return
	}

}
