package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	hash := make(map[string]int)
	ret := make([]string, 0)

	for i := 0; i < len(s); i++ {
		if i+10 <= len(s) {
			tempStr := s[i : i+10]
			count, exists := hash[tempStr]
			if !exists {
				hash[tempStr] = 1
			} else {
				if count == 1 {
					ret = append(ret, tempStr)
				}
				hash[tempStr] += 1
			}
		}
	}

	return ret
}

func main() {
	// fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}
