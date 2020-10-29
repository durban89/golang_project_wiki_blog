package models

import (
	"testing"
)

func Test_Create(t *testing.T) {
	// p := helpers.Page{
	// 	Title: "Test Create",
	// }

	mp := ModelProperty{}

	id, err := mp.Create(InsertValues{})
	if err != nil {
		t.Error(err)
	}

	t.Log(id)
}

func Test_Update(t *testing.T) {
	// p := helpers.Page{
	// 	Title: "Test Update",
	// }

	mp := ModelProperty{}

	effect, err := mp.Update(UpdateValues{}, WhereValues{})
	if err != nil {
		t.Error(err)
	}

	t.Log(effect)
}

func Test_Delete(t *testing.T) {
	mp := ModelProperty{}

	effect, err := mp.Delete(WhereValues{})
	if err != nil {
		t.Error(err)
	}

	t.Log(effect)
}

func TestQueryOne(t *testing.T) {
	mp := ModelProperty{}

	row, err := mp.QueryOne([]string{"user_Id"}, WhereValues{})

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
func Test_Query(t *testing.T) {
	p := ModelProperty{}

	row, err := p.Query([]string{"user_id"}, WhereValues{}, OrderValues{}, 0, 10)

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
