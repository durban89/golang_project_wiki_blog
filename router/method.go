package router

import (
	"net/http"
	"strings"
)

// Method 接口
type Method interface {
	GET(path string, handler func(w http.ResponseWriter, r *http.Request))
	POST(path string, handler func(w http.ResponseWriter, r *http.Request))
	DELETE(path string, handler func(w http.ResponseWriter, r *http.Request))
	PUT(path string, handler func(w http.ResponseWriter, r *http.Request))
	ALL(path string, handler func(w http.ResponseWriter, r *http.Request))
}

// GET 操作
func GET(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if strings.ToLower(r.Method) == "get" {
			handler(w, r)
			return
		}

		http.NotFound(w, r)
		return
	})
}

// POST 操作
func POST(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if strings.ToLower(r.Method) == "post" {
			handler(w, r)
			return
		}

		http.NotFound(w, r)
		return
	})
}

// DELETE 操作
func DELETE(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if strings.ToLower(r.Method) == "delete" {
			handler(w, r)
			return
		}

		http.NotFound(w, r)
		return
	})
}

// PUT 操作
func PUT(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if strings.ToLower(r.Method) == "put" {
			handler(w, r)
			return
		}

		http.NotFound(w, r)
		return
	})
}

// ALL 操作
func ALL(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
		return
	})
}
