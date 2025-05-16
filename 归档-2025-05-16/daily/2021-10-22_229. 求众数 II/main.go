package main

import "fmt"

func majorityElement(nums []int) []int {
	length := len(nums)
	aver := length / 3
	ret := make([]int, 0)
	c := make(map[int]int)
	for _, num := range nums {
		c[num]++
	}

	for key, v := range c {
		if v > aver {
			ret = append(ret, key)
		}
	}

	return ret
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}
