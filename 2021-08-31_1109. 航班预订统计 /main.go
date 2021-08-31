package main

import "fmt"

// func corpFlightBookings(bookings [][]int, n int) []int {
// 	result := make([]int, n)
// 	for _, bookingRecord := range bookings {
// 		for i := bookingRecord[0] - 1; i <= bookingRecord[1] - 1; i++ {
// 			result[i] += bookingRecord[2]
// 		}
// 	}
//
// 	return result
// }

func corpFlightBookings(bookings [][]int, n int) []int {
	result := make([]int, n)
	for _, bookingRecord := range bookings {
		result[bookingRecord[0]-1] += bookingRecord[2]
		if bookingRecord[1] < n {
			result[bookingRecord[1]] -= bookingRecord[2]
		}
	}

	for idx := range result {
		if idx == 0 {
			continue
		}

		result[idx] = result[idx-1] + result[idx]
	}

	return result
}

func main() {
	ret := corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5)
	fmt.Println(ret)
}
