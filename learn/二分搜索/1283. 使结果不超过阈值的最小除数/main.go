package main

import "fmt"

func sumResult(nums []int, t int) int {
	if t == 0 {
		t = 1
	}
	sum := 0
	for _, num := range nums {
		sum += num / t
		if num%t != 0 {
			sum += 1
		}
	}

	return sum
}

func smallestDivisor(nums []int, threshold int) int {
	left, right := 0, 0
	for _, num := range nums {
		if num > right {
			right = num
		}
	}

	for left <= right && right != 0 {
		mid := left + (right-left)>>1
		if sumResult(nums, mid) > threshold {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left == 0 {
		return 1
	}
	return left
}

func main() {
	// fmt.Println(smallestDivisor([]int{1, 2, 5, 9}, 6))
	// fmt.Println(smallestDivisor([]int{2, 3, 5, 7, 11}, 11))
	// fmt.Println(smallestDivisor([]int{21212, 10101, 12121}, 1000000))
	fmt.Println(smallestDivisor([]int{9, 63, 69, 43, 95, 11, 7}, 302))
}
