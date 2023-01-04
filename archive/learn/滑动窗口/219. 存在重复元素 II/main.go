package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= k && (i+j) < len(nums); j++ {
			if nums[i] == nums[i+j] {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	fmt.Println(containsNearbyDuplicate([]int{99, 99}, 2))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 1}, 0))
	fmt.Println(containsNearbyDuplicate([]int{4, 1, 2, 3, 1, 5}, 4))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 2, 3}, 3))
}

/**
1,2,3,1,2,3  k=3
^
^ ^
^   ^
  ^ ^
  ^   ^
    ^ ^
*/
