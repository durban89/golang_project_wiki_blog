package models

import (
	"testing"

	"github.com/durban89/wiki/helpers"
)

func TestCreate(t *testing.T) {
	p := helpers.Page{
		Title: "Test Create",
	}
	id, err := Create(&p)
	if err != nil {
		t.Error(err)
	}

	t.Log(id)
}

func TestUpdate(t *testing.T) {
	p := helpers.Page{
		Title: "Test Update",
	}

	effect, err := Update(&p, 2)
	if err != nil {
		t.Error(err)
	}

	t.Log(effect)
}

func TestDelete(t *testing.T) {
	effect, err := Delete(2)
	if err != nil {
		t.Error(err)
	}

	t.Log(effect)
}

func TestQueryOne(t *testing.T) {
	row, err := QueryOne()

	if err != nil {
		t.Error(err)
	} else {
		t.Log(row)
	}

	if len(row) == 1 {
		t.Log("正确")
	} else {
		t.Error("失败")
	}

	for i, k := range row {
		t.Log(i)
		t.Log(k)
	}
}

// TestQuery 测试获取数据
func TestQuery(t *testing.T) {
	row, err := Query()

	if err != nil {
		t.Error(err)
	} else {
		t.Log(row)
	}

	if len(row) > 0 {
		t.Log("正确")
	} else {
		t.Error("失败")
	}

	for i, k := range row {
		t.Log(i)
		t.Log(k)
	}
}
