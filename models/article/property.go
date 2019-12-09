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
	CategoryID int64
	CreatedAt  string
	UpdatedAt  string
}

func init() {
	Instance = models.ModelProperty{
		TableName: "article",
	}
}
