package article

/*
 * @Author: durban.zhang
 * @Date:   2019-11-29 14:05:25
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2021-02-08 18:51:40
 */

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"

	"wiki/helpers"
	"wiki/models"
	"wiki/models/article"
)

// Save 存储
func Save(w http.ResponseWriter, r *http.Request) {
	session, error := SessionManager.SessionStart(w, r)
	if error != nil {
		http.Error(w, "SessionStart Fail", 403)
		return
	}

	userID := session.Get("user_id")

	session.Get("asd")

	if userID == nil {
		http.Redirect(w, r, "/auth/login", http.StatusFound)
		return
	}

	r.ParseForm()

	var id string

	id = r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	summary := r.FormValue("summary")
	categoryID := r.FormValue("category_id")
	tags := r.FormValue("tags")

	if title == "" || content == "" || summary == "" || categoryID == "" || tags == "" {
		http.Redirect(w, r, helpers.RedirectWithMsg(r, "参数异常"), http.StatusSeeOther)
		return
	}

	// 标题长度限制
	count := utf8.RuneCountInString(title)
	if count > 30 {
		http.Redirect(w, r, helpers.RedirectWithMsg(r, "标题长度太长"), http.StatusSeeOther)
		return
	}

	t := time.Now()
	currentTimeStr := t.Format("2006-01-02 15:04:05")

	if id != "" {
		update := models.UpdateValues{
			"title":       title,
			"content":     content,
			"summary":     summary,
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
			log.Println(err)
			http.Redirect(w, r, helpers.RedirectWithMsg(r, "更新失败"), http.StatusInternalServerError)
			return
		}

		// tags 更新
		updateTag(id, tags)
	} else {
		insert := models.InsertValues{
			"title":       title,
			"content":     content,
			"summary":     summary,
			"category_id": categoryID,
			"author_id":   userID.(string),
			"created_at":  currentTimeStr,
			"updated_at":  currentTimeStr,
		}

		insertID, err := article.Instance.Create(insert)

		saveTag(insertID, tags)

		if err != nil {
			log.Println(err)
			http.Redirect(w, r, helpers.RedirectWithMsg(r, "添加失败"), http.StatusFound)
			return
		}

		id = strconv.FormatInt(insertID, 10)
	}

	if id == "" {
		http.Redirect(w, r, "/articles/view/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/articles/view/?id="+id, http.StatusFound)
	}
}
