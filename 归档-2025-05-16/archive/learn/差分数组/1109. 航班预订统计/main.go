package main

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	df := make([]int, n)
	for _, book := range bookings {
		first, last, count := book[0], book[1], book[2]

		df[first-1] += count
		if last < n {
			df[last] -= count
		}
	}

	// 恢复成标准数组
	for i := range df {
		if i == 0 {
			continue
		}

		df[i] = df[i-1] + df[i]
	}

	return df
}

func main() {
	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
}
