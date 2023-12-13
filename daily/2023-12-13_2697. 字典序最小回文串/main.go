package main

import "fmt"

// func makeSmallestPalindrome(s string) string {
// 	length := len(s)

// 	ret := make([]byte, 0)

// 	for i := 0; i < length/2; i++ {
// 		if s[i] != s[length-i-1] {
// 			if s[i] < s[length-i-1] {
// 				ret = append(ret, s[i])
// 			} else {
// 				ret = append(ret, s[length-i-1])
// 			}
// 		} else {
// 			ret = append(ret, s[i])
// 		}
// 	}

// 	if length%2 != 0 {
// 		ret = append(ret, s[length/2])
// 	}

// 	for i := length/2 - 1; i >= 0; i-- {
// 		ret = append(ret, ret[i])
// 	}

// 	return string(ret)
// }

func makeSmallestPalindrome(s string) string {
	left, right := 0, len(s)-1
	t := []byte(s)

	for left < right {
		if t[left] != t[right] {
			if t[left] > t[right] {
				t[left] = t[right]
			} else {
				t[right] = t[left]
			}
		}
		left++
		right--
	}

	return string(t)
}

func main() {
	ret := makeSmallestPalindrome("seven")
	fmt.Println(ret)
}
