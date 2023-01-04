# 题目

```text
给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/range-sum-query-2d-immutable
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

# 思路

提前计算矩阵和，对于`matrix[i][j]`，其值定义为以`i,j`结尾的矩阵和，

其计算方式为`matrix[i][j] = matrix[i-1][j] + matrix[i][j-1] - matrix[i-1][j-1]`

给定范围的`row1 col1 row2 col2`的矩阵和就为
`matrix[row2][col2]-matrix[row1-1][col2]-matrix[row2][col1-1]+matrix[row1-1][col1-1]`

但是要考虑很多边界情况

# 一点想法
看了标程处理边界方式还是很好的，直接给二维矩阵空出两个多余列，这样就不用处理边界情况了，代码要简单很多