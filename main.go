package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/durban89/wiki/controllers"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// Page 页面结构
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(t string) (*Page, error) {
	filename := "data/" + t + ".txt"
	body, error := ioutil.ReadFile(filename)
	if error != nil {
		return nil, error
	}

	return &Page{
		Title: t,
		Body:  body,
	}, nil
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("无效的页面标题")
	}

	return m[2], nil // 标题是第二个子表达式中
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 这里我们将从Request中提取页面标题，并调用提供的处理程序'fn'
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	p, err := loadPage(title)
// 	if err != nil {
// 		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
// 		return
// 	}

// 	renderTemplate(w, "view", p)
// }

// func editHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	p, err := loadPage(title)
// 	if err != nil {
// 		p = &Page{Title: title}
// 	}

// 	renderTemplate(w, "edit", p)
// }

// func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	body := r.FormValue("body")
// 	p := &Page{
// 		Title: title,
// 		Body:  []byte(body),
// 	}

// 	err := p.save()

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(w, r, "/view/"+title, http.StatusFound)
// }

func main() {
	http.HandleFunc("/view/", makeHandler(controllers.ArticleView))
	http.HandleFunc("/save/", makeHandler(controllers.ArticleSave))
	http.HandleFunc("/edit/", makeHandler(controllers.ArticleEdit))
	log.Fatal(http.ListenAndServe(":8090", nil))
}
