package helpers

import (
	"regexp"
)

// ValidateSliceIntersection 交集
func ValidateSliceIntersection(sourceSlice []string, targetSlice []string) []string {
	var slice3 []string
	for _, slice1 := range sourceSlice {
		for _, slice2 := range targetSlice {
			if slice1 == slice2 {
				slice3 = append(slice3, slice2)
			}
		}
	}
	return slice3
}

// ValidateInArray 元素是否在指定的数组中
func ValidateInArray(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}

	return false
}

// ValidateInt 验证整数
func ValidateInt(str string) bool {
	if m, _ := regexp.MatchString("^[0-9]+$", str); m {
		return true
	}
	return false
}

// ValidateEmail 验证邮箱
func ValidateEmail(str string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, str); m {
		return true
	}

	return false
}

// ValidateChinese 验证中文
func ValidateChinese(str string) bool {
	if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", str); m {
		return true
	}

	return false
}

// ValidateEnglish 验证英文
func ValidateEnglish(str string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", str); m {
		return true
	}

	return false
}
