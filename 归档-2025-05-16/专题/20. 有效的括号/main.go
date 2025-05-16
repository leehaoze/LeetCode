package main

// isValid 返回该字符串是否有效
// 有效指括号必须匹配，例如 "()[]{}" 就是一个有效字符串
func isValid(s string) bool {
	// 当做栈使用
	stack := make([]rune, 0)

	for i := 0; i < len(s); i++ {
		// 字符入栈
		stack = append(stack, (rune)(s[i]))

		// 匹配的字符一起出栈
		for len(stack) >= 2 && match(stack[len(stack)-2], stack[len(stack)-1]) {
			// 出栈
			stack = stack[0 : len(stack)-2]
		}
	}

	return len(stack) == 0
}

func match(a, b rune) bool {
	if a == '[' && b == ']' {
		return true
	}
	if a == '{' && b == '}' {
		return true
	}

	if a == '(' && b == ')' {
		return true
	}

	return false
}
