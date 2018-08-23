package models

import (
	"testing"
)

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
