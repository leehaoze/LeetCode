package main

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	for idx := range words {
		if words[idx][0] != s[idx] {
			return false
		}
	}

	return true
}
