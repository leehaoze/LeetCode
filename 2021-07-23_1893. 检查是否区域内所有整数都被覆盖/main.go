package main

func isCovered(ranges [][]int, left int, right int) bool {
	for i := left; i <= right; i++ {
		in := false
		for _, r := range ranges {
			if i >= r[0] && i <= r[1] {
				in = true
			}
		}

		if !in {
			return false
		}
	}

	return true
}
