package main

import "fmt"

func ship(weights []int, maxWeight int) int {
	count := 0
	ship := 0
	for _, weight := range weights {
		if ship+weight > maxWeight {
			ship = 0
			count++
		}
		ship += weight
	}
	if ship != 0 {
		count++
	}
	return count
}

func shipWithinDays(weights []int, days int) int {
	max, sum := 0, 0
	for _, weight := range weights {
		if weight > max {
			max = weight
		}
		sum += weight
	}

	left, right := max, sum
	for left <= right {
		mid := left + (right-left)>>1
		if ship(weights, mid) > days {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 一定有解
	if left == 0 {
		return 1
	}
	return left
}

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
}
