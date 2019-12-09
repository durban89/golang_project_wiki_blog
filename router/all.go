package router

import (
	"log"
	"net/http"
)

/*
 * @Author: durban.zhang
 * @Date:   2019-12-01 21:21:26
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-02 09:44:25
 */

// ALL Method
func ALL(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
		log.Println("", *r)
		return
	})
}
