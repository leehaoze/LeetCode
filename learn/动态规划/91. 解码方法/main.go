package main

import "fmt"

var exists = map[string]bool{
	"1":  true,
	"2":  true,
	"3":  true,
	"4":  true,
	"5":  true,
	"6":  true,
	"7":  true,
	"8":  true,
	"9":  true,
	"10": true,
	"11": true,
	"12": true,
	"13": true,
	"14": true,
	"15": true,
	"16": true,
	"17": true,
	"18": true,
	"19": true,
	"20": true,
	"21": true,
	"22": true,
	"23": true,
	"24": true,
	"25": true,
	"26": true,
}

// var backup = make(map[string]int)
//
// func numDecodings(s string) int {
// 	if v, exists := backup[s]; exists {
// 		return v
// 	}
// 	length := len(s)
// 	if length == 0 {
// 		return 1
// 	}
//
// 	ret := 0
// 	if length <= 2 {
// 		if exists[s] {
// 			ret += 1
// 		}
// 	}
//
// 	for i := 0; i <= 1 && i < length-1; i++ {
// 		ss := s[0 : i+1]
// 		if exists[ss] {
// 			b := numDecodings(s[i+1:])
// 			if b == 0 {
// 				continue
// 			}
// 			ret += b
// 		}
// 	}
//
// 	backup[s] = ret
// 	return ret
// }

// func numDecodings(s string) int {
// 	length := len(s)
// 	dp := make([][]int, length)
// 	for i := 0; i < length; i++ {
// 		dp[i] = make([]int, length)
// 	}
//
// 	if exists[s[length-1:]] {
// 		dp[length-1][length-1] = 1
// 	}
// 	for i := length - 2; i >= 0; i-- {
// 		if exists[s[i:i+1]] {
// 			dp[i][length-1] = dp[i+1][length-1]
// 		}
//
// 		if exists[s[i:i+2]] {
// 			if i+2 <= length-1 {
// 				dp[i][length-1] += dp[i+2][length-1]
// 			} else {
// 				dp[i][length-1] += 1
// 			}
// 		}
// 	}
//
// 	return dp[0][length-1]
// }

func numDecodings(s string) int {
	length := len(s)
	dp := make([]int, length)
	// if exists[s[length-1:]] {
	// 	dp[length-1] = 1
	// }

	for i := length - 1; i >= 0; i-- {
		if exists[s[i:i+1]] {
			if i == length-1 {
				dp[i] = 1
			} else {
				dp[i] = dp[i+1]
			}
		}

		if i+2 <= length && exists[s[i:i+2]] {
			if i+2 < length {
				dp[i] += dp[i+2]
			} else {
				dp[i] += 1
			}
		}
	}

	return dp[0]
}

func main() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("100"))
}
