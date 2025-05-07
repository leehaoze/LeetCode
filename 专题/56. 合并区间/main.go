package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)

	for _, interval := range intervals {
		if len(merged) == 0 || interval[0] > merged[len(merged)-1][1] {
			merged = append(merged, interval)
		} else {
			// 合并
			merged[len(merged)-1][1] = max(interval[1], merged[len(merged)-1][1])
		}
	}

	return merged
}

func main() {
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {1, 10}}))
}
