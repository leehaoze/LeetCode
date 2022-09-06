package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	preA := 1
	preB := 2
	for i := 3; i <= n; i++ {
		preA, preB = preB, preA+preB
	}
	return preB
}
