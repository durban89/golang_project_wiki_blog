package articlecategory

/*
 * @Author: durban.zhang
 * @Date:   2019-12-31 15:14:55
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-31 15:16:13
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

func init() {
	Instance = models.ModelProperty{
		TableName: "article_category",
	}
}
