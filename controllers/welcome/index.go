package welcome

/*
 * @Author: durban.zhang
 * @Date:   2019-12-12 16:50:31
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-12 17:19:10
 */

import (
	"net/http"

	"github.com/durban89/wiki/helpers"
)

// Index 首页
func Index(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, "index.html", struct {
		Title string
		Data  []string
	}{
		Title: "首页",
		Data:  nil,
	})
	return
}
