package main

import "fmt"

//
// var remoDup = make(map[string]struct{})
// var maxValue = 0
// var maxValueArr = make([][]int, 0)
//
// var memo = make(map[string]int)
//
// func getKey(nums []int) string {
// 	builder := strings.Builder{}
// 	for idx, v := range nums {
// 		builder.WriteString(fmt.Sprintf("idx%d_value%v_", idx, v))
// 	}
// 	fmt.Println(builder.String())
// 	return builder.String()
// }
//
// func countMaxOrSubsets(nums []int) int {
// 	remoDup = make(map[string]struct{})
// 	maxValue = 0
// 	maxValueArr = make([][]int, 0)
// 	memo = make(map[string]int)
//
// 	list(nums, 0)
// 	ret := 0
// 	for i := 0; i < len(maxValueArr); i++ {
// 		key := getKey(maxValueArr[i])
// 		if _, exists := remoDup[key]; !exists {
// 			ret++
// 			remoDup[key] = struct{}{}
// 		}
// 	}
// 	return ret
// }
//
// func cal(nums []int) int {
// 	key := getKey(nums)
// 	if v, exists := memo[key]; exists {
// 		return v
// 	}
// 	// fmt.Println(fmt.Sprintf("Cal %v", nums))
// 	ret := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		ret = ret | nums[i]
// 	}
// 	memo[key] = ret
// 	return ret
// }
//
// func list(nums []int, depth int) [][]int {
// 	if depth == len(nums)-1 {
// 		return [][]int{
// 			{nums[depth]},
// 		}
// 	}
//
// 	subNum := list(nums, depth+1)
// 	newRet := make([][]int, len(subNum)*2)
// 	if len(subNum) > 0 {
// 		for i := 0; i < len(subNum); i++ {
// 			newRet[i] = subNum[i]
// 			newRet[len(subNum)+i] = append(subNum[i], nums[depth])
// 		}
// 	}
// 	newRet = append(newRet, []int{nums[depth]})
//
// 	for _, subArray := range newRet {
// 		ret := cal(subArray)
// 		if ret > maxValue {
// 			maxValue = ret
// 			maxValueArr = make([][]int, 0)
// 			maxValueArr = append(maxValueArr, subArray)
// 		} else if ret == maxValue {
// 			maxValueArr = append(maxValueArr, subArray)
// 		}
// 	}
//
// 	return newRet
// }
func countMaxOrSubsets(nums []int) int {
	maxOr := 0
	count := 0
	for i := 0; i < 1<<len(nums); i++ {
		or := 0
		for j := 0; j < len(nums); j++ {
			if i>>j&1 == 1 {
				or |= nums[j]
			}
		}

		if or > maxOr {
			maxOr = or
			count = 1
		} else if or == maxOr {
			count++
		}

	}
	return count
}

func main() {
	fmt.Println(countMaxOrSubsets([]int{2, 2, 2}))
}
