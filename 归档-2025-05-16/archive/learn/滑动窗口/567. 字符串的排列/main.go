package main

import "fmt"

func checkWindow(s string, start, end int) {
	for i := start; i <= end; i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Print("\n")
}

func checkInclusion(s1 string, s2 string) bool {
	m := make(map[uint8]int)
	for i := 0; i < len(s1); i++ {
		m[s1[i]-'a']++
	}
	containCount := 0

	start, end := 0, 0
	for end < len(s2) {
		c := s2[end] - 'a'
		val, exists := m[c]
		if exists {
			if val > 0 {
				containCount++
			}
			m[c]--
		}
		// checkWindow(s2, start, end)
		for containCount >= len(s1) || ((m[c] < 0 || !exists) && start <= end) {
			if containCount == len(s1) {
				return true
			}

			c = s2[start] - 'a'
			val, e := m[c]
			if e {
				if val >= 0 {
					containCount--
				}
				m[c]++
			}
			start++
		}
		end++
	}

	return false
}

func main() {
	fmt.Println(checkInclusion("abcd", "eiabdfbdca"))
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("hello", "ooolleoooleh"))
	fmt.Println(checkInclusion("adc", "dcda"))
}
