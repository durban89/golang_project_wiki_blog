package article

import (
	"github.com/durban89/wiki/models"
)

var tableName = "article"

// Article 模型
type Article struct {
	models.ModelProperty
}

// Property 属性
type Property struct {
	Autokid string
	Title   string
}
