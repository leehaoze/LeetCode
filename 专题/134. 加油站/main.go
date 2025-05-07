package main

// canCompleteCircuit 计算从哪个加油站出发，可以绕环路形式一周，并返回该加油站位置
// 从 0 位置出发，计算可以抵达的最远位置，如果不能环绕一周，则从最远位置的下一个位置出发
func canCompleteCircuit(gas []int, cost []int) int {
	stationCnt := len(gas)

	for i := 0; i < stationCnt; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		// 最多需要走n过 stationCnt个站点
		for cnt < stationCnt {
			j := (i + cnt) % stationCnt
			sumOfGas += gas[j]
			sumOfCost += cost[j]

			// 这就是能到的最远位置了
			if sumOfCost > sumOfGas {
				break
			}

			cnt++
		}

		// 走过了一圈
		if cnt == stationCnt {
			return i
		}

		// 走不过一圈
		i += cnt + 1
	}

	return -1
}
