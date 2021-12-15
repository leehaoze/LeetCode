package main

// func minWindow(s string, t string) string {
// 	retStr := ""
// 	m := make(map[rune]int)
// 	for _, v := range t {
// 		m[v]++
// 	}
// 	containCount := 0
// 	repeatedCount := 0
//
// 	end := 0
// 	window := strings.Builder{}
// 	for end < len(s) {
// 		// 进行扩容
// 		window.WriteRune(rune(s[end]))
// 		// 进行判断 如果加入窗口的是t中的字符,那么对m进行操作
// 		if val, exists := m[rune(s[end])]; exists {
// 			// 是t中的字符,我们要根据val来操作containCount或者repeatedCount
// 			if val > 0 {
// 				containCount++
// 			} else {
// 				repeatedCount++
// 			}
// 			m[rune(s[end])]--
// 		}
//
// 		// 尝试进行缩容
// 		windowStr := window.String()
// 		_, exists := m[rune(windowStr[0])]
// 		for containCount == len(t) || (repeatedCount > 0 && !exists) {
// 			if containCount == len(t) {
// 				// 这里是包含t中所有的窗口字符串
// 				if retStr == "" || len(retStr) > len(windowStr) {
// 					retStr = windowStr
// 				}
// 			}
//
// 			// 缩小窗口了要 同样要根据剔除的字符来操作containCount或者repeatedCount
// 			c := rune(windowStr[0])
// 			if val, exists := m[c]; exists {
// 				if val >= 0 {
// 					containCount--
// 				} else {
// 					repeatedCount--
// 				}
// 				m[c]++
// 			}
// 			if len(windowStr) > 1 {
// 				windowStr = windowStr[1:]
// 				_, exists = m[rune(windowStr[0])]
// 			} else {
// 				windowStr = ""
// 				exists = true
// 			}
//
// 		}
// 		end++
// 		window.Reset()
// 		window.WriteString(windowStr)
// 	}
//
// 	return retStr
// }
//
// func main() {
// 	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
// 	fmt.Println(minWindow("ABCA", "AAB"))
// 	fmt.Println(minWindow("A", "A"))
// 	fmt.Println(minWindow("A", "AA"))
// 	fmt.Println(minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd"))
// 	fmt.Println(minWindow("aab", "aab"))
// }
