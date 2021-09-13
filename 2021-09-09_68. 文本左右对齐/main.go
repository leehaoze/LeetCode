package main

import (
	"fmt"
	"strings"
)

/**
第一步先将单词分行, 判断方式是 单词 + 1个空格 + 新的单词不超过 maxWidth
*/

func fullJustify(words []string, maxWidth int) []string {
	// step 1 分行
	lengthOfWords := len(words)
	wordInLines := make([][]string, 0, lengthOfWords)
	for i := 0; i < lengthOfWords; i++ {
		// 第一个单词肯定小于maxWidth
		lineWords := make([]string, 0, 2)
		lineWords = append(lineWords, words[i])
		currentLineLength := len(words[i])
		for i+1 < lengthOfWords && currentLineLength+1+len(words[i+1]) <= maxWidth {
			lineWords = append(lineWords, words[i+1])
			currentLineLength += len(words[i+1]) + 1
			i++
		}
		wordInLines = append(wordInLines, lineWords)
	}

	// step 2 每一行填充
	/*
		特殊情况:
			如果这一行只有一个单词,剩余填充为空格
			如果这一行是最后一行,那么应当将空格填充到最后
			对于一个普通行,要尽量将空格均匀分,多余的从左到右挨个填充
	*/
	rowsCount := len(wordInLines)
	ret := make([]string, rowsCount)
	for i := 0; i < rowsCount; i++ {
		// 最后一行
		if i == rowsCount-1 {
			tempString := strings.Builder{}
			for idx, word := range wordInLines[i] {
				if idx == 0 {
					tempString.WriteString(word)
				} else {
					tempString.WriteString(" ")
					tempString.WriteString(word)
				}
			}
			// 填充空格
			blankCount := maxWidth - tempString.Len()
			for i := 0; i < blankCount; i++ {
				tempString.WriteString(" ")
			}
			ret[i] = tempString.String()
		} else if len(wordInLines[i]) == 1 {
			tempString := strings.Builder{}
			tempString.WriteString(wordInLines[i][0])
			// 填充空格
			blankCount := maxWidth - tempString.Len()
			for i := 0; i < blankCount; i++ {
				tempString.WriteString(" ")
			}
			ret[i] = tempString.String()
		} else {
			// 普通行 要计算多余的空格
			// 首先 单词与单词之间是要有一个空格的 可以预先计算这一行的长度
			// 然后与maxLength对比,缺少的长度就是要补足的空格的个数,然后按照单词间隙均分,多余的就依次添加到每个空格里
			length := 0
			for idx, word := range wordInLines[i] {
				length += len(word)
				if idx != 0 {
					length += 1
				}
			}

			blankCount := 1
			extraBlankCount := 0
			if length < maxWidth {
				blankCount += (maxWidth - length) / (len(wordInLines[i]) - 1)
				extraBlankCount += (maxWidth - length) % (len(wordInLines[i]) - 1)
			}

			tempString := strings.Builder{}
			for idx, word := range wordInLines[i] {
				if idx != 0 {
					for i := 0; i < blankCount; i++ {
						tempString.WriteString(" ")
					}

					if extraBlankCount != 0 {
						tempString.WriteString(" ")
						extraBlankCount -= 1
					}
				}
				tempString.WriteString(word)
			}

			ret[i] = tempString.String()

		}

	}
	// for _, s := range ret {
	// 	fmt.Println(s)
	// }
	return ret
}

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	fmt.Println(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
	fmt.Println(fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
}
