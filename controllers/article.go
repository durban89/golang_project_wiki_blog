package controllers

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/durban89/wiki/helpers"
)

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
func ArticleEdit(w http.ResponseWriter, r *http.Request, title string) {
	p, err := helpers.LoadPage(title)
	if err != nil {
		p = &helpers.Page{Title: title}
	}

	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

	p.Token = token
	helpers.RenderTemplate(w, "edit", p)
}
