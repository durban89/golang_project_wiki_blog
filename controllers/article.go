package controllers

import (
	"fmt"
	"net/http"

	"github.com/durban89/wiki/helpers"
)

// ArticleView 查看文章
func ArticleView(w http.ResponseWriter, r *http.Request, title string) {
	p, err := helpers.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	helpers.RenderTemplate(w, "view", p)
}

// ArticleSave 存储文章
func ArticleSave(w http.ResponseWriter, r *http.Request, title string) {
	r.ParseForm()
	fmt.Println("author:", r.Form["author"])

	// if len(r.Form["author"][0]) == 0 {
	// 	fmt.Println("author is empty")
	// 	http.Redirect(w, r, "/view/"+title, http.StatusFound)
	// 	return
	// }

	// if len(r.Form.Get("author")) == 0 {
	// 	fmt.Println("author is empty")
	// 	http.Redirect(w, r, "/view/"+title, http.StatusFound)
	// 	return
	// }

	// getint, geterr := strconv.Atoi(r.Form.Get("author"))
	// if geterr != nil {
	// 	fmt.Println(geterr)
	// 	http.Redirect(w, r, "/view/"+title, http.StatusFound)
	// 	return
	// }

	// if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("author")); !m {
	// 	fmt.Println("非整数")
	// 	http.Redirect(w, r, "/view/"+title, http.StatusFound)
	// 	return
	// }
	// fmt.Println("get author:", r.Form.Get("author"))

	body := r.FormValue("body")
	author := r.FormValue("author")
	method := r.Method

	fmt.Println("author ", author)
	fmt.Println("method: ", method)

	p := &helpers.Page{
		Title: title,
		Body:  []byte(body),
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

	helpers.RenderTemplate(w, "edit", p)
}
