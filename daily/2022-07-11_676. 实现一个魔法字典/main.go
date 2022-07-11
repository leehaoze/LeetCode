package main

import (
	"fmt"
)

type MagicDictionary struct {
	dictionary map[int][]string
}

func Constructor() MagicDictionary {
	return MagicDictionary{dictionary: make(map[int][]string)}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		length := len(word)
		if _, exists := this.dictionary[length]; !exists {
			this.dictionary[length] = make([]string, 0)
		}

		this.dictionary[length] = append(this.dictionary[length], word)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	val, exists := this.dictionary[len(searchWord)]
	if !exists {
		return false
	}

	for _, word := range val {
		if this.exists(word, searchWord) {
			return true
		}
	}

	return false
}

func (this *MagicDictionary) exists(word, searchWord string) bool {
	diffCount := 0
	for idx := range searchWord {
		if searchWord[idx] != word[idx] {
			diffCount++
		}
	}

	return diffCount == 1
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

func main() {
	magicDictionary := Constructor()
	magicDictionary.BuildDict([]string{"hello", "leetcode"})
	fmt.Println(magicDictionary.Search("hello")) // 返回 False
	fmt.Println(magicDictionary.Search("hhllo"))
	fmt.Println(magicDictionary.Search("hell"))
	fmt.Println(magicDictionary.Search("leetcoded"))
}
