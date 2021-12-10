package main

import "fmt"

/*
	1: ()
	2: ()(); (());
	3: ()()(); (()()); ((())); ()(()); (())();
*/

func generateParenthesis(n int) []string {
	return helper("(", n-1, n)
}

func helper(str string, leftCount int, rightCount int) []string {
	// fmt.Println(fmt.Sprintf("Call [%v %v %v]", str, leftCount, rightCount))
	if leftCount == rightCount {
		fmt.Println(fmt.Sprintf("Call [%v %v %v]", str, leftCount, rightCount))
	}
	result := make([]string, 0)
	if leftCount > 0 {
		result = append(result, helper(str+"(", leftCount-1, rightCount)...)
	}

	if rightCount > 0 && rightCount != leftCount {
		result = append(result, helper(str+")", leftCount, rightCount-1)...)
	}

	if leftCount == 0 && rightCount == 0 {
		return []string{str}
	}

	return result
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}
