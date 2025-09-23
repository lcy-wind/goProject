package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//取一组数中只出现一个的数据
	param01 := []int{1, 2, 3, 4, 3, 2, 1}
	fmt.Printf("传入参数为%v, 只出现一个的数据为%v \n", param01, sigleNum(param01))
	// 回文数
	param02 := 121
	fmt.Printf("传入参数为%v, 是回文数吗 %v \n", param02, isPalindrome(param02))
	//有效的括号
	param03 := "("
	fmt.Printf("传入参数为%v, 是有效的括号吗 %v \n", param03, isValid(param03))
	//最长公共前缀
	param04 := []string{"ab", "abc", "abcd"}
	fmt.Printf("传入参数为%v, 最长公共前缀为 %v \n", param04, longestCommonPrefix(param04))
	//加一
	param05 := []int{1, 2, 3, 4}
	fmt.Printf("传入参数为%v, 最大值加1的值为 %v \n", param05, plusOne(param05))
	// 删除有序数组中的重复项
	param06 := []int{1, 1, 2, 3, 4, 5, 5, 6}
	fmt.Printf("传入参数为%v, 顺序取不重复项 %v \n", param06, param06[:removeDuplicates(param06)])
	//合并区间
	param07 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("传入参数为%v, 合并后的区间值 %v \n", param07, merge(param07))
	//两数之和
	param08 := []int{2, 7, 12, 15}
	param08taget := 9
	fmt.Printf("传入参数为%v, 目标值为%v 计算出来下标值为 %v \n", param08, param08taget, twoSum(param08, param08taget))
}

// 定义一个方法，用来找单个数字
func sigleNum(nums []int) int {
	// 创建一个map 用来装key 和key出现的次数
	mapData := make(map[int]int)
	for _, v := range nums {
		mapData[v]++
	}
	// 循环map  查找key出现一次的key
	for k, v := range mapData {
		if v == 1 {
			return k
		}
	}
	return -1
}

// isPalindrome 判断一个整数是否为回文数
func isPalindrome(x int) bool {
	// 负数 和 取余为0的数  不是回文数 直接返回false
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	// 通过x取余 来反转数字
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 如果 x取的值与 反转的数字相等  或者 x的取值 与 反转数字/10相等 也是回文数
	return x == revertedNumber || x == revertedNumber/10
}

// isValid 检查括号字符串是否有效
func isValid(s string) bool {
	// 创建一个映射，存储右括号对应的左括号
	closingToOpening := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 使用切片实现栈
	stack := []rune{}

	// 遍历字符串中的每个字符
	for _, char := range s {
		// 检查当前字符是否是右括号
		if opening, isClosing := closingToOpening[char]; isClosing {
			// 如果是右括号，检查栈是否为空或者栈顶元素是否匹配
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			// 匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，入栈
			stack = append(stack, char)
		}
	}

	// 遍历结束后，栈必须为空才是有效的
	return len(stack) == 0
}

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

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	// 交换元素，使得每个不重复的数字都出现在数组的前面部分。
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1

}

// 合并重叠区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// 按区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 初始化结果切片，将第一个区间加入到结果中
	res := [][]int{intervals[0]}
	// 从第二个区间开始遍历，与结果中的最后一个区间比较，如果当前区间的起始位置小于等于结果中最后一个区间的结束位置，则合并这两个区间；否则将当前区间加入到结果中
	for _, current := range intervals[1:] {
		last := res[len(res)-1]
		if current[0] <= last[1] {
			last[1] = max(last[1], current[1])
		} else {
			res = append(res, current)
		}
	}

	return res
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		sub := target - v
		if j, ok := m[sub]; ok {
			return []int{j, i}
		}
		// 没找到就把当前数存入map
		m[v] = i
	}
	return nil
}
