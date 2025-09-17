package main

import "fmt"

func main() {
	// 题干：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
	// 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
	// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
	// 考察：数字操作、条件判断
	// 题目：判断一个整数是否是回文数

	// 定义一个结构体  输入的值  和 输出的值 输出的值就是正确的单个值
	testData := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 4, 3, 2, 1}, 4},
		{[]int{1, 2, 3, 4, 4, 3, 1}, 2},
		{[]int{1, 2, 3, 4, 4, 2, 1}, 3},
		{[]int{1, 2, 3, 4, 4, 3, 2}, 1},
	}

	for _, v := range testData {
		singNum := sigleNum(v.input)
		fmt.Printf("传入的数据%v, 期望的结果是%v, 计算的结果是%v \n", v.input, v.output, singNum)
	}

}

// 定义一个方法，用来找单个数字
func sigleNum(nums []int) int {
	mapData := make(map[int]int)
	for _, v := range nums {
		mapData[v]++
	}
	for k, v := range mapData {
		if v == 1 {
			return k
		}
	}
	return -1
}
