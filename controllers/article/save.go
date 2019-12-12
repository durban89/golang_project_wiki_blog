package article

/*
 * @Author: durban.zhang
 * @Date:   2019-11-29 14:05:25
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-12 15:54:58
 */

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/durban89/wiki/helpers"
	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
	"github.com/durban89/wiki/models/articletag"
)

// Save 存储
func Save(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var id string

	id = r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	categoryID := r.FormValue("category_id")
	tags := r.FormValue("tags")

	if title == "" || content == "" || categoryID == "" || tags == "" {
		http.Redirect(w, r, helpers.BackWithQuery(r, "参数异常"), http.StatusSeeOther)
		return
	}

	t := time.Now()
	currentTimeStr := t.Format("2006-01-02 15:04:05")

	if id != "" {
		update := models.UpdateValues{
			"title":       title,
			"content":     content,
			"category_id": categoryID,
			"updated_at":  currentTimeStr,
		}

		where := models.WhereValues{
			"autokid": models.WhereCondition{
				Operator: "=",
				Value:    id,
			},
		}

		_, err := article.Instance.Update(update, where)

		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, helpers.BackWithQuery(r, "更新失败"), http.StatusInternalServerError)
			return
		}

		fmt.Println("update to here")

		// tags 更新
		updateTag(id, tags)
	} else {
		insert := models.InsertValues{
			"title":       title,
			"content":     content,
			"category_id": categoryID,
			"created_at":  currentTimeStr,
			"updated_at":  currentTimeStr,
		}

		insertID, err := article.Instance.Create(insert)

		saveTag(insertID, tags)

		if err != nil {
			http.Redirect(w, r, helpers.BackWithQuery(r, "添加失败"), http.StatusSeeOther)
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

func saveTag(articleID int64, tags string) {
	tagsArr := strings.Split(tags, ";")

	t := time.Now()
	currentTimeStr := t.Format("2006-01-02 15:04:05")

	for _, t := range tagsArr {
		if t == "" {
			continue
		}

		var insertTag = models.InsertValues{
			"article_id": strconv.FormatInt(articleID, 10),
			"name":       t,
			"created_at": currentTimeStr,
		}

		_, err := articletag.Instance.Create(insertTag)

		if err != nil {
			panic(err)
		}
	}
}

func updateTag(articleID string, tags string) {
	deleteWhere := models.WhereValues{
		"article_id": models.WhereCondition{
			Operator: "=",
			Value:    articleID,
		},
	}

	_, err := articletag.Instance.Delete(deleteWhere)

	if err != nil {
		panic(err)
	}

	tagsArr := strings.Split(tags, ";")

	t := time.Now()
	currentTimeStr := t.Format("2006-01-02 15:04:05")

	for _, t := range tagsArr {
		if t == "" {
			continue
		}

		var insertTag = models.InsertValues{
			"article_id": articleID,
			"name":       t,
			"created_at": currentTimeStr,
		}

		_, err := articletag.Instance.Create(insertTag)

		if err != nil {
			panic(err)
		}
	}
}
