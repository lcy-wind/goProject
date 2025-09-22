package main

import (
	"fmt"
)

// longestCommonPrefix 取共同前缀
func longestCommonPrefix(strs []string) string {
	// 如果字符串为空则返回空
	if len(strs) == 0 {
		return ""
	}
	// 取第一个字符串作为初始前缀
	for i := 0; i < len(strs[0]); i++ {
		// 先取前缀第一个字符作为前缀，然后依次和后面的字符串比较
		prefix := strs[0][i]
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串长度不够  说明i之前的字符串就是匹配的字符串  或者字符串的相对下标等于截取前缀 就等于公共前缀 直接返回即可
			if i >= len(strs[j]) || strs[j][i] != prefix {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {

	// 	编写一个函数来查找字符串数组中的最长公共前缀。

	// 如果不存在公共前缀，返回空字符串 ""
	testData := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"apple", "app", "application"}, "app"},
		{[]string{"", "a", "ab"}, ""},
		{[]string{"single"}, "single"},
		{[]string{}, ""},
		{[]string{"same", "same", "same"}, "same"},
	}
	for _, v := range testData {
		result := longestCommonPrefix(v.strs)
		fmt.Println("输入：", v.strs, "预期结果", v.expected, "输出：", result)
	}

}
