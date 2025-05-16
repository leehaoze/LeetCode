# 300. 最长递增子序列

子序列的定义是删除数组中任意的元素，求最长严格递增子序列的长度

写了一个位运算的版本，结果人家题目里数组长度很长，位运算直接爆了（毕竟哪有2500位的数据类型）


# 自定向下
定义一个函数 f(nums) => maxLength, minValueOfSubArr
入参是一个数组，出参是最长子序列的长度，以及子序列的最小值

那么对于问题nums, 可得其解为
```go
maxLength, minValueOfSubArr f(nums[1:])
if nums[0] < minValueOfSubArr {
	return maxLength + 1, nums[0]
}
return maxLength, minValueOfSubArr
```

当len(nums) == 1 时
f(nums) = 1, nums[0]


---
事实证明 这个思路是错的
对于这个case[0, 1, 0, 3, 2, 3] 按照这个思路只能得到 [0, 2, 3]这个解
但实际[0, 1, 2, 3]这个解更好

# 动态规划
人家这个dp的定义好
dp[i] 代表的是 以nums[i]这个数结尾的最长递增子序列的长度


