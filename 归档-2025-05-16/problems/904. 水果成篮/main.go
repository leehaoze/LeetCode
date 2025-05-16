package main

import (
	"fmt"
)

func totalFruit(fruits []int) int {
	start, end := 0, 0
	ret := 0

	typeRecord := make(map[int]int)

	for end < len(fruits) {
		typeRecord[fruits[end]]++

		if len(typeRecord) <= 2 && end-start+1 > ret {
			ret = end - start + 1
		}

		for len(typeRecord) > 2 {
			typeRecord[fruits[start]]--
			if typeRecord[fruits[start]] == 0 {
				delete(typeRecord, fruits[start])
			}
			start++
		}
		end++
	}

	return ret
}

func main() {
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}
