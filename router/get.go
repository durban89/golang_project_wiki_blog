package router

import (
	"log"
	"net/http"
)

/*
 * @Author: durban.zhang
 * @Date:   2019-12-01 21:07:14
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-02 09:42:21
 */

// GET Method
func GET(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handler(w, r)
			log.Println("", *r)
			return
		}

		http.NotFound(w, r)
		return
	})
}
