package main

import "fmt"

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

func main() {
	// 测试数据
	cases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
	}

	// 打印结果
	for _, c := range cases {
		fmt.Printf("数组: %v, 目标: %d, 结果: %v\n", c.nums, c.target, twoSum(c.nums, c.target))
	}
}
