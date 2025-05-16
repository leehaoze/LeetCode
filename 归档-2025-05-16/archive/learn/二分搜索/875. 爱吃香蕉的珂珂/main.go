package main

import "fmt"

func eat(piles []int, k int) int {
	time := 0
	for _, pile := range piles {
		time += pile / k
		if pile%k != 0 {
			time += 1
		}
	}

	return time
}

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	// 按照max的速度一定可以在h小时内吃完 但我们是要找最小值
	left, right := 0, max
	for left <= right && right != 0 {
		mid := left + (right-left)>>1
		if eat(piles, mid) <= h {
			right = mid - 1
		} else if eat(piles, mid) > h {
			left = mid + 1
		}
	}

	if left == 0 {
		return 1
	}
	return left
}

func main() {
	fmt.Println(minEatingSpeed([]int{312884470}, 968709470))
}
