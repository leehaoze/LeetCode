# 63.不同路径II

感觉就多了一个障碍物，处理一下这个不通的方案

# 自上向下推

matrix[x-1,y] != 1 && matrix[x,y-1] != 1
    f([x,y]) = f([x-1, y]) + f([x, y-1])
matrix[x-1,y] == 1
    f([x,y]) = f([x, y-1])
matrix([x,y-1]) == 1
    f([x,y]) = f([x-1,y])
matrix[x-1,y] == 1 && matrix[x,y-1] == 1
    f([x,y]) = 0

# 自下向上推

直接跟I一样就可以，只是处理一下障碍物的情况，遇到障碍物要跳过，两个都是障碍物，把当前这个置位0

```go
// init dp first row
	dp[0][0] = 1
	for i := 1; i < column; i++ {
		if dp[0][i-1] == 0 {
			dp[0][i] = 0
		} else if obstacleGrid[0][i] == 0 {
			dp[0][i] = dp[0][i-1]
		} else {
			dp[0][i] = 0
		}
	}
```

初始化第一行和第一列时，其实不用全部走完，只要遇到第一个障碍物，就不用继续初始化了，后边的肯定到不了
