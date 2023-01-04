package main

import "fmt"

func carPooling(trips [][]int, capacity int) bool {
	length := 0
	for i := range trips {
		if trips[i][2] > length {
			length = trips[i][2]
		}
	}

	length = length + 1

	df := make([]int, length)

	for _, trip := range trips {
		people, up, down := trip[0], trip[1], trip[2]

		df[up] += people
		df[down] -= people
	}

	if df[0] > capacity {
		return false
	}
	for i := 1; i < len(df); i++ {
		df[i] = df[i-1] + df[i]
		if df[i] > capacity {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
	fmt.Println(carPooling([][]int{{9, 0, 1}, {3, 3, 7}}, 4))
}
