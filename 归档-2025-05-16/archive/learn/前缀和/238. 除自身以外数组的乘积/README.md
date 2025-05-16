# 238. 除自身以外数组的乘积

```
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

 

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
 

提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。

说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/product-of-array-except-self
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

遍历的时候，首先生成当前的值为[0..i]的数值相乘，再更新ret[0..i]的值乘上当前的nums[i]值，然后超时了
![](http://image.leehaoze.com/20211213152916_DtOE1N_Screenshot.png)
```go
func productExceptSelf(nums []int) []int {
	ret := make([]int, subLen(nums))
	for i := 0; i < subLen(nums); i++ {
		// 初始化 当前值
		ret[i] = 1
		for j := 0; j < i; j++ {
			ret[i] = nums[j] * ret[i]
			ret[j] = ret[j] * nums[i]
		}
	}

	return ret
}
```

---
左右乘积 可以通过正向遍历和反向遍历，计算出每个数的左侧乘积和右侧乘积，然后将每个数的左侧乘积和右侧乘积乘在一起就是解
