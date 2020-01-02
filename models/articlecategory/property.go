package articlecategory

/*
 * @Author: durban.zhang
 * @Date:   2019-12-31 15:14:55
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 17:21:50
 */

import "github.com/durban89/wiki/models"

// Instance 实例
var Instance models.ModelProperty

// Property 属性
type Property struct {
	Autokid   int64
	Name      string
	Desc      string
	CreatedAt string
	UpdateAt  string
}

// QueryField 查询字段
func QueryField() models.SelectValues {
	p := Property{}

	return models.SelectValues{
		"autokid":    &p.Autokid,
		"name":       &p.Name,
		"desc":       &p.Desc,
		"created_at": &p.CreatedAt,
		"updated_at": &p.UpdateAt,
	}
}

func init() {
	Instance = models.ModelProperty{
		TableName:  "article_category",
		QueryFiled: QueryField(),
	}
}
