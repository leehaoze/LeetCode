package main

import "strconv"

func fizzBuzz(n int) []string {
	ret := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ret[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			ret[i-1] = "Fizz"
		} else if i%5 == 0 {
			ret[i-1] = "Buzz"
		} else {
			ret[i-1] = strconv.FormatInt(int64(i), 10)
		}
	}
	return ret
}
