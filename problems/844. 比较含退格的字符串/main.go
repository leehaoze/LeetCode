package main

import (
	"fmt"
)

func backspaceCompare(s string, t string) bool {
	// return trace(s) == trace(t)
	sp := len(s) - 1
	tp := len(t) - 1

	for sp >= 0 || tp >= 0 {
		skip := 0
		for sp >= 0 {
			if s[sp] == '#' {
				skip++
				sp--
			} else if skip > 0 {
				sp--
				skip--
			} else {
				break
			}
		}

		skip = 0
		for tp >= 0 {
			if t[tp] == '#' {
				skip++
				tp--
			} else if skip > 0 {
				tp--
				skip--
			} else {
				break
			}
		}

		if sp < 0 && tp < 0 {
			return true
		}

		if sp < 0 || tp < 0 || s[sp] != t[tp] {
			return false
		}

		sp--
		tp--

	}

	return true
}

// func trace(s string) string {
// 	left, right := 0, 0
// 	data := make([]uint8, len(s))
// 	for right < len(data) {
// 		if s[right] == '#' {
// 			if left > 0 {
// 				left--
// 			}
// 			right++
// 			continue
// 		}
//
// 		data[left] = s[right]
// 		left++
// 		right++
// 	}
// 	// fmt.Println(string(data[:left]))
// 	return string(data[:left])
// }

func main() {
	fmt.Println(backspaceCompare("y#fo##f", "y#f#o##f"))
}
