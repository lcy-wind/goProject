package main

import "fmt"

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

func main() {
	// 测试案例
	testCases := []int{121, -121, 10, 12321, 12345, 0}
	for _, num := range testCases {
		result := isPalindrome(num)
		var resultStr string
		if result {
			resultStr = "是"
		} else {
			resultStr = "不是"
		}
		fmt.Printf("数字 %d 是回文数吗？ %v\n", num, resultStr)
	}

}
