package main

import (
	"fmt"
	"sort"
	"strings"
)

// 输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
// 输出："the cat was rat by the bat"
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/replace-words
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func replaceWords(dictionary []string, sentence string) string {
	sort.SliceStable(dictionary, func(i, j int) bool {
		if len(dictionary[i]) == len(dictionary[j]) {
			return dictionary[i] < dictionary[j]
		}
		return len(dictionary[i]) < len(dictionary[j])
	})

	tokens := strings.Split(sentence, " ")
	for idx := range tokens {
		tokens[idx] = replace(tokens[idx], &dictionary)
	}

	return strings.Join(tokens, " ")
}

func replace(token string, dictionary *[]string) string {
	for _, s := range *dictionary {
		if strings.HasPrefix(token, s) {
			return s
		}
	}

	return token
}

func main() {
	dictionary := []string{"a", "b", "c"}
	fmt.Println(replaceWords(dictionary, "aadsfasf absbs bbab cadsfafs"))
}
