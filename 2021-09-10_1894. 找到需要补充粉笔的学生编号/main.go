package main

import "fmt"

func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, i := range chalk {
		sum += i
	}

	left := k % sum
	for idx, i := range chalk {
		if i <= left {
			left -= i
		} else {
			return idx
		}
	}
	return 0
}

func main() {
	fmt.Println(chalkReplacer([]int{5,1,5}, 22))
	fmt.Println(chalkReplacer([]int{3,4,1,2}, 25))
}
