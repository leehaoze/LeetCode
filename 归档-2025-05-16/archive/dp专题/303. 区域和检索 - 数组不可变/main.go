package main

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	return NumArray{preSum: preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
