# 解1

先找到旋转点，通过旋转点将nums分为两个部分，[4,5,6,1,2,3]可以分为[4,5,6]和[1,2,3]，且第一部分的数组值全部大于第二部分。 通过判断target值是否大于第一部分的最小值nums[0]
，可以判断出target落入了哪个区间，对那个区间进行二分即可。 第一发wrong掉是因为没有考虑旋转点为0，即没旋转的情况，加了一个特殊判断处理才A掉

# 解2-错误思路

对于nums进行二分时，落入的[left, right]区间有这么几种可能性:

- 区间是严格递增的，区间完全在左半部分，或者完全在右半部分，标志为nums[left] < nums[right]，这种情况继续正常二分即可
- 区间不是严格递增的，区间横跨了两个部分，标志位nums[left] > nums[right]，因为左半区间的任意值均大于右半区间的任意值。这种情况需要判断target是大于nums[mid]还是小于nums[mid]
    - 如果是大于nums[mid]，反而应该将right缩小到mid-1继续查找
    - 如果是小于nums[mid]，则应该将left扩大到mid+1

# 解2
对于有旋转的nums数组,[4,5,6,1,2,3]，通过mid切分为两部分后，一定会存在一个有序的部分，可以分开讨论情况：
- mid落入左侧区间，即左侧区间是有序的
  target只要不在左侧区间[left, mid-1]这个范围，那么就要去右侧区间搜索
- mid落入右侧区间，即右侧区间是有序的
  target只要不在右侧区间[mid+1, right]这个范围，那么就去左侧区间搜索

```go
func search(nums []int, target int) int {
	left, right := 0, subLen(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		fmt.Printf("left: %v  right: %v mid:%v\n", left, right, mid)
		if nums[mid] == target {
			return mid
		}
		// 左侧有序
		if nums[left] < nums[mid-1] {
			if target >= nums[left] && target <= nums[mid-1] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右侧有序
			if target >= nums[mid+1] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
```
写到这里因为边界问题导致在右侧有序分支，if target >= nums[mid+1]这里越界了，看了标程

发现人家判断有序的方式很简单，直接判断nums[mid] 与 nums[0]的关系即可
nums[mid] > nums[0]，则左侧有序，否则右侧有序

写了一版代码后，还是会在target >= nums[mid+1] 这个地方越界。看了标程，人家没有用到mid+1
我这里写的是target >= nums[mid+1] 其实等价于 target > nums[mid]，用后者就没有越界问题了