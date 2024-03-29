package util

import "strings"

// ContainString 判断一个字符串是否在一个字符串切片中.
// target: 目标字符串
// slice: 字符串切片
// mustMatchCase: 是否必须大小写匹配
func ContainString(target string, slice []string, mustMatchCase bool) bool {
	for _, str := range slice {
		if mustMatchCase {
			if str == target {
				return true
			}
		} else {
			if strings.ToLower(str) == strings.ToLower(target) {
				return true
			}
		}
	}
	return false
}

func ContainKey(target string, dict map[string]string, mustMatchCase bool) bool {
	for key := range dict {
		if mustMatchCase {
			if key == target {
				return true
			}
		} else {
			if strings.ToLower(key) == strings.ToLower(target) {
				return true
			}
		}
	}
	return false
}

// IsSubSlice 判断一个字符串切片是否是另一个字符串切片的子集.
// subSlice: 子集
// slice: 父集
// mustMatchCase: 是否必须大小写匹配
func IsSubSlice(subSlice []string, slice []string, mustMatchCase bool) bool {
	for _, str := range subSlice {
		if !ContainString(str, slice, mustMatchCase) {
			return false
		}
	}
	return true
}
