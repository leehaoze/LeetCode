package main

// import (
// 	"fmt"
// 	"strings"
// )
//
// func contain(s string, t string) bool {
// 	m := make(map[rune]int)
// 	for _, c := range s {
// 		m[c]++
// 	}
//
// 	for _, c := range t {
// 		if val, exists := m[c]; exists {
// 			if val == 0 {
// 				return false
// 			}
// 			m[c]--
// 		} else {
// 			return false
// 		}
// 	}
//
// 	return true
// }
//
// func minWindow(s string, t string) string {
// 	retStr := ""
//
// 	end := 0
// 	window := strings.Builder{}
// 	for end < len(s) {
// 		window.WriteRune(rune(s[end]))
// 		currentWindow := window.String()
// 		for contain(currentWindow, t) {
// 			if retStr == "" || len(retStr) > len(currentWindow) {
// 				retStr = currentWindow
// 			}
//
// 			// 缩小窗口
// 			currentWindow = currentWindow[1:]
// 		}
// 		window.Reset()
// 		window.WriteString(currentWindow)
// 		end++
// 	}
//
// 	return retStr
// }
//
// func main() {
// 	// fmt.Println(contain("ABCA", "BAAA"))
// 	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
// 	fmt.Println(minWindow("ABCA", "AAB"))
// 	fmt.Println(minWindow("A", "A"))
// 	fmt.Println(minWindow("A", "AA"))
// 	fmt.Println(minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd"))
// }
