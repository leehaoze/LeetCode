package main

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, pos := range position {
		if pos%2 == 0 {
			odd++
		} else {
			even++
		}
	}

	if odd < even {
		return odd
	}

	return even
}
