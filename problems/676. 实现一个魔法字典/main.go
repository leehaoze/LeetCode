package main

import "fmt"

// type MagicDictionary struct {
// 	dict map[int][]string
// }

// func Constructor() MagicDictionary {
// 	return MagicDictionary{
// 		dict: make(map[int][]string),
// 	}
// }

// func (this *MagicDictionary) BuildDict(dictionary []string) {
// 	for _, word := range dictionary {
// 		if _, exists := this.dict[len(word)]; !exists {
// 			this.dict[len(word)] = make([]string, 0)
// 		}

// 		this.dict[len(word)] = append(this.dict[len(word)], word)
// 	}
// }

// func (this *MagicDictionary) Search(searchWord string) bool {
// 	length := len(searchWord)

// 	wordsToMatch := this.dict[length]

// 	for _, word := range wordsToMatch {
// 		diffCount := 0
// 		for i := 0; i < length; i++ {
// 			if word[i] != searchWord[i] {
// 				diffCount++
// 			}

// 			if diffCount > 1 {
// 				break
// 			}
// 		}

// 		if diffCount == 1 {
// 			return true
// 		}
// 	}

// 	return false
// }

type trieTree struct {
	children map[rune]*trieTree
	isEnd    bool
}

func NewTrieTree() *trieTree {
	return &trieTree{
		children: make(map[rune]*trieTree),
		isEnd:    false,
	}
}

func (t *trieTree) insert(word string) {
	node := t
	for _, c := range word {
		if _, exists := node.children[c]; !exists {
			node.children[c] = NewTrieTree()
		}

		node = node.children[c]
	}
	node.isEnd = true
}

func (t *trieTree) search(node *trieTree, word string, modified bool) bool {
	if word == "" {
		// 如果word已经没有了，则要看是否也到了字典树的终点
		// 必须修改过一次字符
		return modified && node.isEnd
	}

	if node.children[rune(word[0])] != nil && t.search(node.children[rune(word[0])], word[1:], modified) {
		return true
	}

	if modified {
		return false
	}

	// 修改一次 挨个查询
	for c, child := range node.children {
		if c != rune(word[0]) && t.search(child, word[1:], true) {
			return true
		}
	}

	return false
}

type MagicDictionary struct {
	*trieTree
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		trieTree: NewTrieTree(),
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		this.insert(word)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return this.search(this.trieTree, searchWord, false)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

func main() {
	obj := Constructor()

	obj.BuildDict([]string{"hello", "hallo", "leetcode"})
	fmt.Println(obj.Search("hell"))
}
