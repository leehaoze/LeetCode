package main

import "fmt"

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	// 0 代表前 0 项的和
	prefix[0] = 0
	for i, num := range nums {
		prefix[i+1] = prefix[i] + num
	}
	// fmt.Println(prefix)
	return NumArray{prefix: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right+1] - this.prefix[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main() {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(numArray.SumRange(0, 2))
	fmt.Println(numArray.SumRange(2, 5))
	fmt.Println(numArray.SumRange(0, 5))
}
