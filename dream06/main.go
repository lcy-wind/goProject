package main

import "fmt"

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

func main() {
	testData := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 5},
		{},
		{5},
		{1, 2, 3, 4, 5},
	}

	for _, nums := range testData {
		// 复制原数组用于输出展示
		original := make([]int, len(nums))
		copy(original, nums)

		k := removeDuplicates(nums)
		fmt.Printf("原始数组: %v, 处理后前%d个元素: %v\n", original, k, nums[:k])
	}
}
