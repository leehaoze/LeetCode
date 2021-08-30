package main

import "fmt"

var table [101]int
var inited = false

func getMaximumGenerated(n int) int {
	if !inited {
		initTable()
	}

	return table[n]
}

func initTable() {
	table[1] = 1
	for i := 0; i <= 100; i++ {
		if 2 <= 2*i && 2*i <= 100 {
			table[2*i] = table[i]
		}

		if 2 <= 2*i+1 && 2*i+1 <= 100 {
			table[2 * i + 1] = table[i] + table[i + 1]
		}
	}

	max := 0
	for i := 0; i <= 100; i++ {
		if max > table[i] {
			table[i] = max
		} else {
			max = table[i]
		}
	}



	inited = true
}

func main() {
	fmt.Println(getMaximumGenerated(7))
	fmt.Println(getMaximumGenerated(2))
	fmt.Println(getMaximumGenerated(3))
}