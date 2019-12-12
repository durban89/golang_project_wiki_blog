package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-02 10:55:36
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-12 16:43:12
 */

import (
	"net/http"

	"github.com/durban89/wiki/helpers"
	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
)

// Delete 删除操作
func Delete(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")

	if id == "" {
		http.NotFound(w, r)
		return
	}

	articleID := ""

	var selectValues = models.SelectValues{
		"autokid": &articleID,
	}

	var whereValues = models.WhereValues{
		"autokid": models.WhereCondition{
			Operator: "=",
			Value:    id,
		},
	}

	err := article.Instance.QueryOne(selectValues, whereValues)

	if err != nil || articleID == "" {
		http.Redirect(
			w,
			r,
			helpers.BackWithQuery(r, "操作失败"),
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
