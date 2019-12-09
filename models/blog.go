package models

// var tableName = "blog"

// BlogModel 模型
type BlogModel struct {
	ModelProperty
}

// BlogProperty 属性
type BlogProperty struct {
	Autokid string
	Title   string
}
