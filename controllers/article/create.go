package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:55:18
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-02 18:46:11
 */

import (
	"net/http"

	"github.com/durban89/wiki/helpers"
)

// Create 文件
func Create(w http.ResponseWriter, r *http.Request) {
	// 视图渲染
	helpers.Render(w, "create.html", nil)

	return
}
