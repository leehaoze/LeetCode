package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	// 1 3 5 7 8 10 12 是31
	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	dates := strings.Split(date, "-")
	year, _ := strconv.ParseInt(dates[0], 10, 64)
	if (year%100 == 0 && year%400 == 0) || year%4 == 0 {
		days[2]++
	}

	month, _ := strconv.ParseInt(dates[1], 10, 64)
	day, _ := strconv.ParseInt(dates[2], 10, 64)
	count := 0
	for i := int64(0); i < month; i++ {
		count += days[i]
	}
	count += int(day)
	return count
}

func main() {
	fmt.Println(dayOfYear("2003-03-01"))
}
