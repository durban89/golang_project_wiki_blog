package views

/*
 * @Author: durban.zhang
 * @Date:   2020-01-02 18:14:57
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 18:15:45
 */

import "strconv"

// ToInt 整数
func ToInt(value interface{}) int {
	switch v := value.(type) {
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0
		}
		return i
	case int:
		return v
	default:
		return 0
	}
}
