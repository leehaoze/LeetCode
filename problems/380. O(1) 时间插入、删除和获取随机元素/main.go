package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	// hash val => 数组中的索引
	hash map[int]int
	// data 存储所有数据，用于随机返回
	data []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hash: make(map[int]int),
		data: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	// 先判断是否存在
	if _, exists := this.hash[val]; exists {
		return false
	}

	// 开始插入操作
	length := len(this.hash)

	if len(this.data) <= length {
		this.data = append(this.data, val)
	} else {
		this.data[length] = val
	}

	this.hash[val] = length
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// fmt.Println(fmt.Sprintf("%#v", this.hash))

	// 删除元素
	if _, exists := this.hash[val]; !exists {
		return false
	}

	// 找到下标位置
	idx := this.hash[val]

	// 删除此元素，将他与最后一个元素交换位置即可
	this.data[idx], this.data[len(this.hash)-1] = this.data[len(this.hash)-1], this.data[idx]

	// fmt.Println(fmt.Sprintf("idx: %v, this.data[idx]: %v, this.data[len(this.hash)-1]: %v", idx, this.data[idx], this.data[len(this.hash)-1]))

	// 更新hash表
	this.hash[this.data[idx]] = idx

	// 删除hash表
	delete(this.hash, this.data[len(this.hash)-1])

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.hash))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	obj := Constructor()
	// obj.Remove(0)
	// obj.Remove(0)
	obj.Insert(0)
	obj.Remove(0)
	fmt.Println(obj.Insert(0))
}
