package main

import (
	"fmt"
	"strconv"
	"strings"
)

func evaluate(expression string) int {
	return cal(expression, make(map[string]int))
}

// var deep = 0

func cal(expression string, scope map[string]int) int {
	// deep++
	// for i := 0; i < deep; i++ {
	// 	fmt.Print("----")
	// }
	// defer func() {
	// 	deep--
	// }()
	// fmt.Println("expression: ", expression)
	token := split(expression)
	// for i := 0; i < deep; i++ {
	// 	fmt.Print("----")
	// }
	// fmt.Println(fmt.Sprintf("tokens: %#v", token))
	if token[0] == "let" {
		// let 从idx=1开始是赋值表达式，最后一个是运算表达式
		for i := 1; i < len(token)-2; i += 2 {
			scope[token[i]] = cal(token[i+1], clone(scope))
		}
		// fmt.Println("expr: ", token[len(token)-1])
		return cal(token[len(token)-1], clone(scope))
	} else if token[0] == "mult" {
		a := cal(token[1], clone(scope))
		b := cal(token[2], clone(scope))
		// fmt.Println(fmt.Sprintf("mult %v * %v", a, b))
		return a * b
	} else if token[0] == "add" {
		a := cal(token[1], clone(scope))
		b := cal(token[2], clone(scope))
		// fmt.Println(fmt.Sprintf("add %v + %v", a, b))
		return a + b
	} else {
		t := strings.Replace(token[0], ")", "", -1)
		// 对于 数字 直接return
		val, err := strconv.ParseInt(t, 10, 64)
		if err == nil {
			return int(val)
		}
		// 对于 x y 这种表达式 直接 return值
		return scope[t]
	}
}

func split(expr string) []string {
	// let v e v e v e ... expr
	// add e1 e2
	// mult e1 e2
	if strings.Contains(expr, "(") {
		// 去掉首尾括号
		expr = expr[1 : len(expr)-1]
		tokens := strings.Split(expr, " ")
		ret := make([]string, 0)
		temp := make([]string, 0)

		subExpr := false
		leftCount := 0
		for _, token := range tokens {
			if strings.Contains(token, "(") {
				leftCount += strings.Count(token, "(")
				subExpr = true
				temp = append(temp, token)
				continue
			}

			if strings.Contains(token, ")") {
				leftCount -= strings.Count(token, ")")
				if leftCount == 0 {
					subExpr = false
					temp = append(temp, token)
					ret = append(ret, strings.Join(temp, " "))
					temp = make([]string, 0)
					continue
				}
			}

			if subExpr {
				temp = append(temp, token)
			} else {
				ret = append(ret, token)
			}
		}

		if len(temp) > 0 {
			ret = append(ret, strings.Join(temp, " "))
		}

		return ret
	} else {
		return strings.Split(expr, " ")
	}
}

func clone(scope map[string]int) map[string]int {
	targetMap := make(map[string]int)
	for key, val := range scope {
		targetMap[key] = val
	}
	return targetMap
}

func main() {
	// fmt.Println(evaluate("(let x 2 y (add x 1) (mult x y))"))
	// fmt.Println(evaluate("(let x 2 (mult x (let x 3 y 4 (add x y))))"))
	fmt.Println(evaluate("(let x 2 (add (let x 3 (let x 4 x)) x))"))
}
