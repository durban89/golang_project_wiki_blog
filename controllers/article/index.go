package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:53:13
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-31 17:51:03
 */

import (
	"fmt"
	"net/http"
	"time"

	"github.com/durban89/wiki/helpers"
	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
)

// Index 默认页面
func Index(w http.ResponseWriter, r *http.Request) {
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

	// var articleModel article.Article

	var autokid int64
	var title string
	selectField := models.SelectValues{
		"autokid": &autokid,
		"title":   &title,
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

	helpers.Render(w, "article/index.html", struct {
		Data   []models.SelectResult
		Cookie string
	}{
		Data:   qr,
		Cookie: siteName,
	})

	return
}
