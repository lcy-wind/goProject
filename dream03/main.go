package main

import "fmt"

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

func main() {
	// 	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
	// 有效字符串需满足：
	// 左括号必须用相同类型的右括号闭合。
	// 左括号必须以正确的顺序闭合。
	// 每个右括号都有一个对应的相同类型的左括号。
	testData := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"(", false},
		{")", false},
	}

	// 执行测试并输出结果
	for _, tc := range testData {
		result := isValid(tc.s)
		fmt.Printf("输入: %q, 预期: %v, 结果: %v\n",
			tc.s, tc.expected, result)
	}

}
