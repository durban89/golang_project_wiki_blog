package welcome

/*
 * @Author: durban.zhang
 * @Date:   2019-12-12 16:50:31
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 14:38:20
 */

import (
	"net/http"

	"wiki/views"
)

// Index 首页
func Index(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "welcome/index.html", struct {
		Title string
		Data  []string
	}{
		Title: "首页",
		Data:  nil,
	})
	return
}
