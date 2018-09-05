package models

import (
	"fmt"
	"strings"
)

// Model 模型
type Model interface {
	QueryOne()
	Update()
}

// WhereCondition where条件
type WhereCondition struct {
	Operator string
	Value    string
}

// WhereValues where条件值
type WhereValues map[string]WhereCondition

// UpdateValues update条件值
type UpdateValues map[string]string

// SelectValues select条件值
type SelectValues []string

// MergeWhere 合并where条件
func (w WhereValues) MergeWhere() string {
	where := []string{}
	for k, v := range w {
		if v.Operator == "" {
			s := fmt.Sprintf("%s = %s", k, v.Value)
			where = append(where, s)
		} else {
			s := fmt.Sprintf("%s %s %s", k, v.Operator, v.Value)
			where = append(where, s)
		}
	}

	return strings.Join(where, " AND ")
}

// MergeUpdate 合并update条件
func (u UpdateValues) MergeUpdate() string {
	update := []string{}
	for k, v := range u {
		s := fmt.Sprintf("%s = '%s'", k, v)
		update = append(update, s)
	}

	return strings.Join(update, ", ")
}

// MergeSelect  合并select条件
func (s SelectValues) MergeSelect() string {
	return strings.Join(s, ", ")
}
