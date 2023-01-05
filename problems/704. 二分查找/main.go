package _04__二分查找

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 左闭右闭
	for left <= right {           // 区间还有值，因为是左闭右闭区间，因此left == right的时候也是有值的，因此不能退出循环
		mid := left + (right-left)/2 // 求中间值的坐标，右移两位快速除2
		if nums[mid] == target {     // 找到目标 返回下标
			return mid
		} else if nums[mid] > target { // 目标在左半区间，因此收缩right
			right = mid - 1 // 因为是左闭右闭区间，因此right缩小到mid-1，不缩到mid的原因是mid已经被排除掉了
		} else {
			left = mid + 1
		}
	}

	return -1
}
