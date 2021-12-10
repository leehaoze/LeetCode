package main

import (
	"fmt"
	"sort"
)

func isSubStr(target, s string) bool {
	idx := 0
	idxOfTarget := 0
	for idx < len(s) {
		// find s[idx] in target
		find := false
		for idxOfTarget < len(target) {
			idxOfTarget++
			if s[idx] == target[idxOfTarget-1] {
				find = true
				break
			}
		}
		if !find {
			return false
		}
		idx++
	}
	return true
}

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) != len(dictionary[j]) {
			return len(dictionary[i]) > len(dictionary[j])
		} else {
			for idx := range dictionary[i] {
				if dictionary[i][idx] != dictionary[j][idx] {
					return dictionary[i][idx] < dictionary[j][idx]
				}
			}
		}
		return true
	})

	for _, word := range dictionary {
		i := 0
		for j := range s {
			if s[j] == word[i] {
				i++
				if i == len(word) {
					return word
				}
			}
		}
		// if isSubStr(s, word) {
		// 	return word
		// }
	}

	return ""
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))
}
