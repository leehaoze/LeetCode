package main

import "fmt"

/*
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000


I 在 V 和 X的左侧，表示 4和9
X 在 L 和 C左边，表示40和90
C 在 D 和 M左变，表示400和900
*/

func romanToInt(s string) int {
	ret := 0

	charToInt := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if s[i] == 'I' && i+1 < len(s) && (s[i+1] == 'V' || s[i+1] == 'X') {
			ret += (charToInt[s[i+1]] - charToInt[s[i]])
			i++
			continue
		}

		if s[i] == 'X' && i+1 < len(s) && (s[i+1] == 'L' || s[i+1] == 'C') {
			ret += (charToInt[s[i+1]] - charToInt[s[i]])
			i++
			continue
		}

		if s[i] == 'C' && i+1 < len(s) && (s[i+1] == 'D' || s[i+1] == 'M') {
			ret += (charToInt[s[i+1]] - charToInt[s[i]])
			i++
			continue
		}

		ret += charToInt[s[i]]

	}

	return ret
}

func main() {
	fmt.Println(romanToInt("IV"))
}
