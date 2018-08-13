package helpers

import "io/ioutil"

// Page 页面结构
type Page struct {
	Title string
	Body  []byte
}

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
