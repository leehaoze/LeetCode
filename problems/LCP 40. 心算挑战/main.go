package main

import (
	"slices"
)

func maxmiumScore(cards []int, cnt int) int {
	slices.Sort(cards)
	slices.Reverse(cards)

	tmpSum := 0
	minOdd := -1
	minEven := -1

	for i := 0; i < cnt; i++ {
		tmpSum += cards[i]

		if cards[i]%2 == 1 {
			// 这个地方不需要判断是否比 minOdd 小，因为数组是有序的，越往后越小，直接覆盖就行
			minOdd = cards[i]
		} else {
			minEven = cards[i]
		}
	}

	if tmpSum%2 == 0 {
		return tmpSum
	}

	// 分情况讨论
	ans := 0
	for i := cnt; i < len(cards); i++ {
		if cards[i]%2 == 1 {
			if minEven != -1 {
				ans = max(ans, tmpSum-minEven+cards[i])
			}
		} else {
			if minOdd != -1 {
				ans = max(ans, tmpSum-minOdd+cards[i])
			}
		}
	}

	return ans
}
