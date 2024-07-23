package main

import (
	"sort"
)

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)

	maxRet := -1
	sum := 0
	for idx, val := range beans {
		tmp := val * (len(beans) - idx)
		sum += val

		if tmp > maxRet {
			maxRet = tmp
		}
	}

	return int64(sum - maxRet)
}
