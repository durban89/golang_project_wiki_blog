package helpers

import (
	"html/template"
	"io/ioutil"
)

// Page 页面结构
type Page struct {
	ID     int64
	Title  string
	Body   []byte
	Script string
	Html   template.HTML
	Token  string
}

// Save 存储数据
func (p *Page) Save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// LoadPage 记载文章
func LoadPage(t string) (*Page, error) {
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
