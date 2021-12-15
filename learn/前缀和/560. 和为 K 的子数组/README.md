给你一个整数数组 nums 和一个整数k ，请你统计并返回该数组中和为k的连续子数组的个数。

示例 1：

输入：nums = [1,1,1], k = 2 输出：2 示例 2：

输入：nums = [1,2,3], k = 3 输出：2

---
可以使用前缀和数组，计算出前缀和数组后，遍历这个前缀和数组来求解

---

```go
for i := 0; i < len(preSum); i++ {
    for j := i + 1; j < len(preSum); j++ {
        if preSum[j]-preSum[i] == k {
            count++
        }
    }
}
```

遍历前缀和数组的过程如上，这是求解 i～len(preSum)区间有没有合法的j 满足 `preSum[j] - preSum[i] == k`。 我们也可以修改这个区间为 `0~i`，这样遍历的是经过的区间

```go
for i := 1; i < len(preSum); i++ {
    for j := i-1; j >= 0; j -- {
		if preSum[i] - preSum[j] == k {
			count++
        }
    }
}
```

我们寻找合法的 `j` 的过程，其实在寻找一个下标j(这个j的范围是0~i)对应的preSum值是`preSum[i] - k`
那我们可以在生成preSum数组时，直接生成一个`[前缀和值A] => [前缀和值为A的个数]`这样形式的map，同时遍历每个i时，我们只需要看一下map中`map[preSum[i]-k]`的Val即可

```go
count := 0
valMap := make(map[int]int)
valMap[0] = 1
for i := 0; i < len(nums); i++ {
	preSum[i+1] = preSum[i] + nums[i]
	valMap[preSum[i+1]]++
	count += valMap[perSum[i+1]-k]
}
```

至于初始化map是的`valMap[0]=1`，想一下这个map的定义，前缀和值A的数量，什么情况下我们需要找一个`preSum[j] = 0`的前缀和
是`preSum[i] - k = 0`，也就是`preSum[i]`刚好与我们要找的k一致，那么此时我们不再需要找另外的数据项来凑出k来，`preSum[i]`本身就是满足和为k的一个子数组，因此其val为1