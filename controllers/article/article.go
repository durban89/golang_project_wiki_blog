package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-30 17:49:57
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 10:07:36
 */

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/durban89/wiki/models"
	"github.com/durban89/wiki/models/article"
	"github.com/durban89/wiki/models/articlecategory"
	"github.com/durban89/wiki/models/articletag"
	"github.com/durban89/wiki/session"
)

// SessionManager 初始化session
var SessionManager *session.Manager

func init() {
	var err error
	SessionManager, err = session.GetManager("memory", "sessionid", 3600)
	if err != nil {
		log.Println(err)
		return
	}

	go SessionManager.SessionGC()
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

func getArticleCategory(categoryID string) map[string]string {
	var id string
	var name string

	category := make(map[string]string)

	selectValue := models.SelectValues{
		"autokid": &id,
		"name":    &name,
	}

	whereValue := models.WhereValues{
		"autokid": models.WhereCondition{
			Operator: "=",
			Value:    categoryID,
		},
	}

	err := articlecategory.Instance.QueryOne(selectValue, whereValue)

	if err != nil {
		return nil
	}

	category["name"] = name
	category["id"] = id

	log.Println(category)

	return category
}

func getArticleCategories() []models.SelectResult {
	var id int64
	var name string

	selectValue := models.SelectValues{
		"autokid": &id,
		"name":    &name,
	}

	whereValue := models.WhereValues{}
	orderValue := models.OrderValues{
		"autokid": models.OrderCondition{
			OrderBy: "DESC",
		},
	}

	category, err := articlecategory.Instance.Query(selectValue, whereValue, orderValue, 0, 1000)

	if err != nil {
		return nil
	}

	log.Println(category)

	return category
}

func getArticleTag(w http.ResponseWriter, r *http.Request, id string) []string {
	// tag query
	var tagName string
	var tagsArr = []string{}

	selectTagField := models.SelectValues{
		"name": &tagName,
	}

	whereTag := models.WhereValues{
		"article_id": models.WhereCondition{
			Operator: "=",
			Value:    id,
		},
	}

	order := models.OrderValues{
		"autokid": models.OrderCondition{
			OrderBy: "DESC",
		},
	}

	tags, err := articletag.Instance.Query(selectTagField, whereTag, order, 0, 100)

	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return tagsArr
	}

	for _, v := range tags {
		tagsArr = append(tagsArr, v["name"].(string))
	}

	return tagsArr
}

func getAuthor(authorID string) (map[string]string, error) {
	author := map[string]string{}

	author["username"] = "dpzhang"
	author["email"] = "admin@126.com"
	author["authorID"] = authorID

	return author, nil

	var username string
	var email string

	selectField := models.SelectValues{
		"username": &username,
		"email":    &email,
	}

	where := models.WhereValues{
		"autokid": models.WhereCondition{
			Operator: "=",
			Value:    authorID,
		},
	}

	err := article.Instance.QueryOne(selectField, where)

	if err != nil {
		return nil, err
	}

	author["username"] = username
	author["email"] = email
	author["authorID"] = authorID

	return author, nil
}
