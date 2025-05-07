package main

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	// key 是排序后的单词
	retMap := make(map[string][]string)

	for _, str := range strs {
		key := sortStr(str)
		retMap[key] = append(retMap[key], str)
	}

	ret := make([][]string, 0)
	for _, val := range retMap {
		ret = append(ret, val)
	}

	return ret
}

func sortStr(str string) string {
	data := make([]rune, 0)
	for _, c := range str {
		data = append(data, c)
	}

	slices.Sort(data)

	ret := ""
	for _, c := range data {
		ret += string(c)
	}

	return ret
}
