package main

func shipWithinDays(weights []int, D int) int {
	// 查找解的范围
	left := weights[0]
	right := 0
	for _, val := range weights {
		right += val
		if left < val {
			left = val
		}
	}

	// maxValue <= 解 <= sum
	for left < right {
		mid := (left + right) / 2
		if tryShip(mid, weights) <= D {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func tryShip(shipWeight int, weights []int) (needDay int) {
	sum := 0
	needDay = 1
	for i := range weights {
		if sum+weights[i] <= shipWeight {
			sum = sum + weights[i]
		} else {
			sum = weights[i]
			needDay += 1
		}
	}

	return needDay
}
