package router

/*
 * @Author: durban.zhang
 * @Date:   2019-12-01 21:27:36
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-01 21:28:11
 */

import "net/http"

// Method 接口
type Method interface {
	GET(path string, handler func(w http.ResponseWriter, r *http.Request))
	POST(path string, handler func(w http.ResponseWriter, r *http.Request))
	DELETE(path string, handler func(w http.ResponseWriter, r *http.Request))
	PUT(path string, handler func(w http.ResponseWriter, r *http.Request))
	ALL(path string, handler func(w http.ResponseWriter, r *http.Request))
}
