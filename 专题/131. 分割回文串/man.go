package main

func partition(s string) [][]string {
	
}

func isPalindrome(s string, start, end int) bool {
	i, j := start, end
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
