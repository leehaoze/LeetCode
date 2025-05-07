package main

func minWindow(s string, t string) string {
	// 预处理t
	ori := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	cnt := make(map[byte]int)

	// 滑动窗口检测函数
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}

		return true
	}

	// 滑动窗口，左开右闭
	left, right := 0, 0
	ret := ""
	for right < len(s) {
		cnt[s[right]]++
		right++

		for check() && left <= right {
			if ret == "" || right-left < len(ret) {
				ret = s[left:right]
			}

			cnt[s[left]]--
			left++
		}

	}

	return ret
}
