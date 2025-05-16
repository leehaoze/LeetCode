package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	words := make([]string, 0)

	var word string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(word) == 0 {
				continue
			}

			words = append(words, word)
			word = ""

		} else {
			word += string(s[i])
		}
	}

	if len(word) != 0 {
		words = append(words, word)
		word = ""
	}

	slices.Reverse(words)
	// fmt.Println(fmt.Sprintf("%#v", words))
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("  the sky is blue"))
}
