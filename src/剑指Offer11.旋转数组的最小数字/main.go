package main

//func minArray(numbers []int) int {
//	for i := 1; i < len(numbers); i++ {
//		if numbers[i] < numbers[i-1] {
//			return numbers[i]
//		}
//	}
//
//	return numbers[0]
//}

func minArray(numbers []int) int {
	var i, j int = 0, len(numbers) - 1

	for i < j {
		m := (i + j) / 2
		if numbers[m] > numbers[j] {
			i = m + 1
		} else if numbers[m] < numbers[j] {
			j = m
		} else {
			j--
		}
	}

	return numbers[i]
}
