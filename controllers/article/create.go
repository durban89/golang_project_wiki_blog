package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:55:18
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-30 18:04:20
 */

import (
	"fmt"
	"net/http"
	"time"

	"github.com/durban89/wiki/helpers"
)

// Create 文件
func Create(w http.ResponseWriter, r *http.Request) {
	session, err := SessionManager.SessionStart(w, r)
	if err != nil {
		http.Error(w, "Session启动失败", 500)
		return
	}

	userID := session.Get("user_id")

	fmt.Println(userID)

	if userID == nil {
		http.Error(w, "无权限访问", 403)
		return
	}

	msg := r.URL.Query().Get("err_msg")

	t := time.Now()
	timeStr := t.Format("2006-01-02 03:04:05")

	// 视图渲染
	helpers.Render(w, "article/create.html", struct {
		Msg  string
		Time string
	}{
		Msg:  msg,
		Time: timeStr,
	})

	return
}
