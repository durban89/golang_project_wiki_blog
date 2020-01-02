package views

/*
 * @Author: durban.zhang
 * @Date:   2020-01-02 18:15:59
 * @Last Modified by:   durban.zhang
 * @Last Modified time: 2020-01-02 18:16:19
 */

import "strconv"

// ToInt64 64整数
func ToInt64(value interface{}) int64 {
	switch v := value.(type) {
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0
		}
		return i
	case int64:
		return v
	default:
		return 0
	}
}
