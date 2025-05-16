package main

import (
	"fmt"
	"strings"
)

/**
5F3Z-2e-9-w
*/
func licenseKeyFormatting(s string, k int) string {
	originStr := strings.Builder{}
	for i := range s {
		if s[i] != '-' {
			originStr.WriteByte(s[i])
		}
	}

	length := originStr.Len()
	first := length % k
	if first == 0 {
		first = k
	}

	retStr := strings.Builder{}
	groupElemCount := first
	count := 0
	for _, s := range originStr.String() {
		if count == groupElemCount {
			retStr.WriteByte('-')
			groupElemCount = k
			count = 0
		}

		if s >= 'a' && s <= 'z' {
			retStr.WriteByte(byte(s + 'A' - 'a'))
		} else {
			retStr.WriteByte(byte(s))
		}
		count += 1
	}

	return retStr.String()
}

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
}
