package main

import (
	"fmt"
	"sort"
)

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

func main() {
	// 测试用例
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {2, 3}},
		{{5, 5}},
		{},
	}

	// 打印结果
	for _, intervals := range testCases {
		fmt.Printf("输入: %v\n", intervals)
		fmt.Printf("合并后: %v\n\n", merge(intervals))
	}
}
