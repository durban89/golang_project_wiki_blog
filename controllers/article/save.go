package article

/*
 * @Author: durban.zhang
 * @Date:   2019-11-29 14:05:25
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-11-29 14:15:19
 */

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
)

// Save 存储
func Save(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var id string

	id = r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	category := r.FormValue("category")

	if title == "" || content == "" || category == "" {
		fmt.Println("参数异常")
	}

	if id != "" {
		update := models.UpdateValues{
			"title":   title,
			"content": content,
		}
		where := models.WhereValues{
			"autokid": models.WhereCondition{
				Operator: "=",
				Value:    id,
			},
		}

		_, err := article.Instance.Update(update, where)

		if err != nil {
			http.Redirect(w, r, "/articles/error", http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Println("crete")
		insert := models.InsertValues{
			"title":   title,
			"content": content,
		}

		insertID, err := article.Instance.Create(insert)

		fmt.Println(insertID)

		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/articles/error", http.StatusSeeOther)
			return
		}

		id = strconv.FormatInt(insertID, 10)
	}

	if id == "" {
		http.Redirect(w, r, "/articles/view/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/articles/view/?id="+id, http.StatusSeeOther)
	}
}
