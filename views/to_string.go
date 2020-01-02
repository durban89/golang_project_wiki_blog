package views

/*
 * @Author: durban.zhang
 * @Date:   2020-01-02 18:14:50
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 18:15:33
 */

import "strconv"

// ToString 字符串
func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	default:
		return ""
	}
}
