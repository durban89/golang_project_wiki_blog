package article

/*
 * @Author: durban.zhang
 * @Date:   2019-12-30 17:49:57
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2021-02-08 18:05:06
 */

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"wiki/models"
	"wiki/models/article"
	ArticleCategory "wiki/models/article/category"
	ArticleTag "wiki/models/article/tag"
	"wiki/session"

	// memory session provider
	_ "wiki/session/providers/memory"
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

		_, err := ArticleTag.Instance.Create(insertTag)

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

	_, err := ArticleTag.Instance.Delete(deleteWhere)

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

		_, err := ArticleTag.Instance.Create(insertTag)

		if err != nil {
			panic(err)
		}
	}
}

func getArticleCategory(categoryID string) models.SelectResult {
	whereValue := models.WhereValues{
		"autokid": models.WhereCondition{
			Operator: "=",
			Value:    categoryID,
		},
	}

	category, err := ArticleCategory.Instance.QueryOne([]string{
		"autokid", "name",
	}, whereValue)

	if err != nil {
		log.Println(err)
		return nil
	}

	return category
}

func getArticleCategories() []models.SelectResult {
	whereValue := models.WhereValues{}
	orderValue := models.OrderValues{
		"autokid": models.OrderCondition{
			OrderBy: "DESC",
		},
	}

	category, err := ArticleCategory.Instance.Query([]string{
		"autokid",
		"name",
	}, whereValue, orderValue, 0, 1000)

	if err != nil {
		return nil
	}

	return category
}

func getArticleTag(w http.ResponseWriter, r *http.Request, id string) []string {
	// tag query
	var tagsArr = []string{}

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

	tags, err := ArticleTag.Instance.Query([]string{
		"name",
	}, whereTag, order, 0, 100)

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

func getAuthor(authorID string) (models.SelectResult, error) {
	author := make(models.SelectResult)

	author["username"] = "dpzhang"
	author["email"] = "admin@126.com"
	author["authorID"] = authorID

	return author, nil

	selectField := []string{
		"username",
		"email",
	}

	where := models.WhereValues{
		"autokid": models.WhereCondition{
			Operator: "=",
			Value:    authorID,
		},
	}

	author, err := article.Instance.QueryOne(selectField, where)

	if err != nil {
		return nil, err
	}

	// author["username"] = username
	// author["email"] = email
	author["authorID"] = authorID

	return author, nil
}
