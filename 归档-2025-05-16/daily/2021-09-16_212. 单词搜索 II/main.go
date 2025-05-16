package main

func wordContain(points []byte, word string) (bool, string) {
	for i := range points {
		if points[i] != word[i] {
			return false, ""
		}
	}

	if len(points) == len(word) {
		return true, word
	}

	return true, ""
}

func wordMatch(points []byte, words []string) (bool, []string) {
	match := false
	ret := make([]string, 0)
	for i := range words {
		if matched, word := wordContain(points, words[i]); matched {
			match = true
			if len(word) != 0 {
				ret = append(ret, word)
			}
		}
	}

	return match, ret
}

func findWords(board [][]byte, words []string) []string {
	step := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			// 没有这个点开头的单词就跳过了
			if matched, _ := wordMatch([]byte{board[i][j]}, words); !matched {
				continue
			}

			// 初始化节点
			stack := make([]byte, 0)
			stack = append(stack, board[i][j])

			length := 1
			for len(step) != 0 {
				// 从这个点往下走
				point := step[length-1]
				// 开始移动

			}
		}
	}
}
