package main

import "fmt"

func generateParenthesis(n int) []string {
	dp := make([]string, 1)
	dp[0] = "()"

	for i := 0; i < n-1; i++ {
		ret := make([]string, 0)
		for _, solve := range dp {
			ret = append(ret, fmt.Sprintf("(%s)", solve))
			ret = append(ret, fmt.Sprintf("()%s", solve))
			s := fmt.Sprintf("%s()", solve)
			// if ret[len(ret)-1] != s {
			// 	ret = append(ret, fmt.Sprintf("()%s", solve))
			// }
			exists := false
			for i := 0; i < len(ret); i++ {
				if ret[i] == s {
					exists = true
				}
			}

			if !exists {
				ret = append(ret, s)
			}
		}
		dp = ret
	}

	return dp
}

func main() {
	fmt.Println(generateParenthesis(4))
	// fmt.Println("()()()" == "()()()")
}
