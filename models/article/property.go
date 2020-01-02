package article

import (
	"github.com/durban89/wiki/models"
)

// Instance 实例
var Instance models.ModelProperty

// Property 属性
type Property struct {
	Autokid    int64
	Title      string
	Content    string
	Summary    string
	CategoryID int64
	AuthorID   int64
	CreatedAt  string
	UpdatedAt  string
}

// QueryField 查询字段
func QueryField() models.SelectValues {
	property := Property{}

	return models.SelectValues{
		"autokid":     &property.Autokid,
		"title":       &property.Title,
		"content":     &property.Content,
		"summary":     &property.Summary,
		"category_id": &property.CategoryID,
		"author_id":   &property.AuthorID,
		"created_at":  &property.CreatedAt,
		"updated_at":  &property.UpdatedAt,
	}
}

func init() {
	Instance = models.ModelProperty{
		TableName:  "article",
		QueryFiled: QueryField(),
	}
}
