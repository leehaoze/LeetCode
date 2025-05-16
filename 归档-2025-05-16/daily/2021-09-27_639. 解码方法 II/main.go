package main

import "fmt"

var countMap = map[string]int{
	"*": 9,
	"1": 1,
	"2": 1,
	"3": 1,
	"4": 1,
	"5": 1,
	"6": 1,
	"7": 1,
	"8": 1,
	"9": 1,
}

func numDecodings(s string) int {
	if count, exists := countMap[s]; exists {
		return count
	}

	if len(s) == 1 {
		if s[0] == '*' {
			return 9
		}

		return 1
	}

	if len(s) == 2 {
		// 第一位可以是 1~2 或者*
		// 第二位可以使 1~6 或者*
		if ((s[0] >= '1' && s[0] <= '2') || (s[0] == '*')) && ((s[1] >= '1' && s[1] <= '6') || (s[1] == '*')) {
			decodeA := 1
			if s[0] == '*' {
				decodeA = 2
			}
			decodeB := 1
			if s[1] == '*' {
				if s[0] == '1' {
					decodeB = 9
				} else {
					decodeB = 6
				}
			}
			if s[0] == '*' && s[1] == '*' {
				decodeA = 1
				decodeB = 26
			}


			return numDecodings(string(s[0])) * numDecodings(string(s[1])) + decodeA*decodeB
		} else {
			return numDecodings(string(s[0])) * numDecodings(string(s[1]))
		}
	}

	return numDecodings(s[:2]) + numDecodings(s[1:])
}

/**
A 1
B 2
C 3
D 4
E 5
F 6
G 7
H 8
I 9
J 10
K 11
L 12
M 13
N 14
O 15
P 16
Q 17
R 18
S 19
T 20
U 21
V 22
W 23
X 24
Y 25
Z 26

1 -> 'A' 有一种解
2 -> 'B' 一种解
12 -> 'AB' 'L'
*/

func main() {
	// fmt.Println(numDecodings("1*"))
	fmt.Println(numDecodings("2*"))
}
