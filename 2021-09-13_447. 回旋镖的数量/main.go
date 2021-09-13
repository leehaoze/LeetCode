package main

import (
	"fmt"
)

/**
两两比较点之间的距离,距离相同的放到一个数组里,然后计算数组里的点数量就可以
*/

func calDistance(a []int, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

func numberOfBoomerangs(points [][]int) int {

	ans := 0

	for _, p := range points {
		sameDistance := make(map[int]int)
		for _, q := range points {
			sameDistance[calDistance(p, q)] ++
		}

		for _, m := range sameDistance {
			ans += m * (m - 1)
		}

	}

	return ans
}

func main() {
	fmt.Println(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
}
