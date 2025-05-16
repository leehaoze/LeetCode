package main

import "fmt"

/*
[2,1] => [1,2] | [2,1]
[2,1],[3,2] => [?,2,?]
[2,1],[3,4],[3,2] => [?,2,?] + [?,3,?]
[1,2,3]|[3,2,1] + [2,3,4]|[4,3,2]
----------------------------------------
[4,-2] => [4,-2] | [-2,4]
[4,-2],[1,4] => [1,4,-2] | [-2,4,1]
[4,-2],[1,4],[-3,1] => [1,4,-2] | [-2,4,1] + [4,1,-3]|[-3,4,1]
*/

func restoreArray(adjacentPairs [][]int) []int {
	appearCount := make(map[int]int)
	for _, pair := range adjacentPairs {
		appearCount[pair[0]]++
		appearCount[pair[1]]++
	}

	first := 0
	for i := range appearCount {
		if appearCount[i] == 1 {
			first = i
			break
		}
	}

	ret := make([]int, 0)
	ret = append(ret, first)

	used := make(map[int]bool)

	find := func(head int) int {
		for idx, pair := range adjacentPairs {
			if !used[idx] {
				if pair[0] == head || pair[1] == head {
					used[idx] = true
					if pair[0] == head {
						return pair[1]
					} else {
						return pair[0]
					}
				}
			}
		}

		return -1
	}

	for len(ret) <= len(adjacentPairs)  {

		ret = append(ret, find(ret[len(ret)-1]))
	}

	return ret
}

func main() {
	fmt.Println(restoreArray([][]int{
		{2, 1},
		{3, 4},
		{3, 2},
	}))

	fmt.Println(restoreArray([][]int{
		{4, -2},
		{1, 4},
		{-3, 1},
	}))

	fmt.Println(restoreArray([][]int{
		{100000, -100000},
	}))

	fmt.Println(restoreArray([][]int{
		{4, -10},
		{-1, 3},
		{4, -3},
		{-3, 3},
	}))
}
