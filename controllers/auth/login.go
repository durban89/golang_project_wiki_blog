/*
* @Author: durban.zhang
* @Date:   2019-12-30 10:15:18
* @Last Modified by:   durban.zhang
* @Last Modified time: 2020-01-02 14:38:16
 */

package auth

import (
	"fmt"
	"log"
	"net/http"

	"wiki/helpers"
	"wiki/views"
)

// Login Page
func Login(w http.ResponseWriter, r *http.Request) {
	session, err := SessionManager.SessionStart(w, r)
	userID := session.Get("user_id")

	if r.Method == "POST" {
		r.ParseForm()

		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" {
			helpers.RedirectWithMsg(r, "参数不能为空")
			return
		}

		log.Println(email)
		log.Println(password)

		if email == "admin@126.com" && password == "123" {

			if err == nil {
				session.Set("user_id", "1")
			}

			http.Redirect(w, r, fmt.Sprintf("/?msg=登录成功"), http.StatusFound)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/auth/login/?msg=账号或密码错误"), http.StatusFound)
		return
	}

	if userID != nil {
		http.Redirect(w, r, "/?msg=已经登录", http.StatusFound)
		return
	}

	views.Render(w, "auth/login.html", nil)
	return
}
