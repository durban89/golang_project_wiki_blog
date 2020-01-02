package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:53:13
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 14:37:53
 */

import (
	"fmt"
	"net/http"
	"time"

	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
	"github.com/durban89/wiki/views"
)

// Index 默认页面
func Index(w http.ResponseWriter, r *http.Request) {
	session, error := SessionManager.SessionStart(w, r)
	if error != nil {
		http.Error(w, "SessionStart Fail", 403)
		return
	}

	userID := session.Get("user_id")

	var siteName string
	cookie, err := r.Cookie("site_name_cookie")

	if err != nil {
		expired := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{
			Name:    "site_name_cookie",
			Value:   "gowhich_cookie",
			Expires: expired,
		}

		http.SetCookie(w, &cookie)
	} else {
		siteName = cookie.Value
	}

	var autokid int64
	var title string
	var authorID string
	var created string

	selectField := models.SelectValues{
		"autokid":    &autokid,
		"title":      &title,
		"author_id":  &authorID,
		"created_at": &created,
	}

	where := models.WhereValues{}
	order := models.OrderValues{
		"autokid": models.OrderCondition{
			OrderBy: "DESC",
		},
	}

	qr, err := article.Instance.Query(selectField, where, order, 0, 10)

	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}

	views.Render(w, "article/index.html", struct {
		Data   []models.SelectResult
		UserID interface{}
		Cookie string
	}{
		Data:   qr,
		UserID: userID,
		Cookie: siteName,
	})

	return
}
