package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:55:18
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-09 16:52:58
 */

import (
	"net/http"
	"time"

	"github.com/durban89/wiki/helpers"
)

// Create 文件
func Create(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("err_msg")

	t := time.Now()
	timeStr := t.Format("2006-01-02 03:04:05")

	// 视图渲染
	helpers.Render(w, "create.html", struct {
		Msg  string
		Time string
	}{
		Msg:  msg,
		Time: timeStr,
	})

	return
}
