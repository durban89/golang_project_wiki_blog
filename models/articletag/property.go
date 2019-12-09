package articletag

/*
 * @Author: durban.zhang
 * @Date:   2019-12-09 16:45:36
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2019-12-09 19:13:18
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

func init() {
	Instance = models.ModelProperty{
		TableName: "article_tag",
	}
}
