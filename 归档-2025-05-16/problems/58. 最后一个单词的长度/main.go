package main

import "fmt"


func lengthOfLastWord(s string) int {
	var word string
	var ret int
	for i := 0; i < len(s); i++ {
		c := s[i]

		if c == ' ' && len(word) == 0 {
			continue
		}

		if c == ' ' && len(word) != 0 {
			ret = len(word)
			word = ""
			continue
		}

		word += string(c)
	}

	if len(word) != 0 {
		ret = len(word)
	}
	return ret
}


func main() {
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}