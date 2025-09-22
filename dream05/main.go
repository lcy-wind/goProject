package main

import (
	"fmt"
	"strconv"
)

// longestCommonPrefix 取共同前缀
func plusOne(strs []int) []int {
	if len(strs) == 0 {
		return []int{}
	}
	var sumStr string
	// 转换为字符串，然后转换为整数
	for _, v := range strs {
		sumStr += strconv.Itoa(v)
	}
	// 转换为整数
	sum, err := strconv.ParseInt(sumStr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	sum++
	// 转换为字符串
	resultStr := strconv.FormatInt(sum, 10)
	// 转换回整数数组
	result := make([]int, len(resultStr))
	for i, c := range resultStr {
		result[i], _ = strconv.Atoi(string(c))
	}
	return result
}

func main() {
	// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
	testData := []struct {
		strs     []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{6, 2, 3}, []int{6, 2, 4}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{}, []int{}},
	}
	for _, v := range testData {
		result := plusOne(v.strs)
		fmt.Printf("输入: %v, 预计输出: %v, 实际输出: %v\n", v.strs, v.expected, result)

	}

}
