package main

func nthUglyNumber(n int) int {
	start := [...]int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 20, 25}
	if n <= len(start) {
		return start[n-1]
	}

}
