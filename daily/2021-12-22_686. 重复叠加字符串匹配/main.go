package main

import "fmt"

func repeatedStringMatch(a string, b string) int {
	m := make(map[uint8][]int)
	length := len(a)
	for i := 0; i < length; i++ {
		c := a[i] - 'a'
		if _, exists := m[c]; !exists {
			m[c] = make([]int, 0)
		}
		m[c] = append(m[c], i)
	}

	var possibleStart []int
	possibleStart, exists := m[b[0]-'a']
	if !exists {
		return -1
	}

	for _, start := range possibleStart {
		// fmt.Println(fmt.Sprintf("start:%v", start))
		var i int
		count := 0
		pos := start
		for i = 0; i < len(b); i++ {
			pos = start + i
			if pos >= length {
				pos = pos % length
			}
			if pos == 0 {
				count++
			}
			// fmt.Println(fmt.Sprintf("i is %v", i))
			// fmt.Println(fmt.Sprintf("compare b[%v]=%c, a[%v]=%c", i, b[i], pos, a[pos]))
			if b[i] != a[pos] {
				break
			}
		}
		if pos != length-1 {
			count++
		}
		if i == len(b) {
			return count
		}
	}

	return -1
}

func main() {
	// fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("aa", "aaaaa"))
}
