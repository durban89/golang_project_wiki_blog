package articletag

/*
 * @Author: durban.zhang
 * @Date:   2019-12-09 16:45:36
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 16:32:09
 */

import "github.com/durban89/wiki/models"

// Instance 实例
var Instance models.ModelProperty

// Property 属性
type Property struct {
	Autokid   int64
	ArticleID int64
	Name      string
	CreatedAt string
}

// QueryField 查询字段
func QueryField() models.SelectValues {
	p := Property{}

	return models.SelectValues{
		"autokid":    &p.Autokid,
		"article_id": &p.ArticleID,
		"name":       &p.Name,
		"created_at": &p.CreatedAt,
	}
}

func init() {
	Instance = models.ModelProperty{
		TableName:  "article_tag",
		QueryFiled: QueryField(),
	}
}
