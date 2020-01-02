package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:55:36
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 17:09:30
 */

import (
	"net/http"

	"github.com/durban89/wiki/helpers"
	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
)

// Delete 删除操作
func Delete(w http.ResponseWriter, r *http.Request) {
	session, error := SessionManager.SessionStart(w, r)
	if error != nil {
		http.Error(w, "SessionStart Fail", 403)
		return
	}

	userID := session.Get("user_id")

	if userID == nil {
		http.Redirect(w, r, "/auth/login", http.StatusFound)
		return
	}

	var id = r.URL.Query().Get("id")

	if id == "" {
		http.NotFound(w, r)
		return
	}

	// articleID := ""

	// var selectValues = models.SelectValues{
	// 	"autokid": &articleID,
	// }

	var selectValues = []string{
		"autokid",
	}

	var whereValues = models.WhereValues{
		"autokid": models.WhereCondition{
			Operator: "=",
			Value:    id,
		},
	}

	_, err := article.Instance.QueryOne(selectValues, whereValues)

	if err != nil {
		http.Redirect(
			w,
			r,
			helpers.RedirectWithMsg(r, "操作失败"),
			http.StatusSeeOther,
		)
		return
	}

	article.Instance.Delete(whereValues)

	http.Redirect(
		w,
		r,
		"/articles/",
		http.StatusSeeOther,
	)

	return
}
