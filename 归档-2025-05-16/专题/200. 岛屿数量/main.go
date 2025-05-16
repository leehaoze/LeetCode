package main

var opX = [4]int{1, 0, -1, 0} // 右、下、左、上
var opY = [4]int{0, 1, 0, -1}

// numsIslands 判断二维网格中，岛屿的数量
func numIslands(grid [][]byte) int {
	// 初始化标识数组
	// visit := make([][]bool, len(grid))
	// for i := 0; i < len(grid); i++ {
	// 	visit[i] = make([]bool, len(grid[i]))
	// }

	var ret int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 已经访问过
			// if visit[i][j] {
			// 	continue
			// }

			// 不是岛屿
			if grid[i][j] == '0' {
				continue
			}

			// 遍历网格，以当前网格为起点，出发
			ret++
			toVisit := make([]int, 0)
			toVisit = append(toVisit, i, j)
			grid[i][j] = '0'
			for len(toVisit) != 0 {
				// visit[toVisit[0][0]][toVisit[0][1]] = true
				// 沉默岛屿

				for op := 0; op < 4; op++ {
					newPosX := toVisit[0] + opX[op]
					newPosY := toVisit[1] + opY[op]

					// 判断是否越界
					if isLegal(grid, newPosX, newPosY) && grid[newPosX][newPosY] == '1' {
						toVisit = append(toVisit, newPosX, newPosY)
						grid[newPosX][newPosY] = '0'
					}
				}

				if len(toVisit) >= 2 {
					toVisit = toVisit[2:]
				}
			}
		}
	}

	// for i := 0; i < len(visit); i++ {
	// 	fmt.Println(visit[i])
	// }

	return ret
}

func isLegal(grid [][]byte, i, j int) bool {
	if i < 0 || i >= len(grid) {
		return false
	}

	if j < 0 || j >= len(grid[0]) {
		return false
	}

	return true
}
