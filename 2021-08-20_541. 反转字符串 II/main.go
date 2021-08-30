package main

import (
	"fmt"
	"strings"
)

/**
abcas
"hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl"
39
*/
func reverseStr(s string, k int) string {
	length := len(s)
	ret := strings.Builder{}
	ret.Grow(length)
	for i := 0; i < length; {
		if i+k-1 >= length {
			for j := length-1; j >= i ; j-- {
				ret.WriteByte(s[j])
			}
		} else {
			for j := i + k - 1; j >= i; j-- {
				ret.WriteByte(s[j])
			}
		}

		for j := i + k; j < i+k*2 && j < length; j++ {
			ret.WriteByte(s[j])
		}

		i = i + k*2
	}

	return ret.String()
}

func main() {
	fmt.Println(reverseStr("abcdav", 2))
	fmt.Println(reverseStr("abcd", 2))
	fmt.Println(reverseStr("abcd", 3))
	fmt.Println(reverseStr("abcd", 4))
	fmt.Println(reverseStr("abcdefga", 3))
	// fmt.Println("hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl")
	fmt.Println(reverseStr("hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl", 39))
	// fmt.Println("fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi")
}
