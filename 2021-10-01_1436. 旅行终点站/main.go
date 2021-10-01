package main

import "fmt"

func destCity(paths [][]string) string {
	firstMap := make(map[string]struct{})
	for _, path := range paths {
		firstMap[path[0]] = struct{}{}
	}

	for _, path := range paths {
		if _, exists := firstMap[path[1]]; !exists {
			return path[1]
		}
	}

	return ""
}

func main() {
	fmt.Println(destCity([][]string{{"London","New York"},{"New York","Lima"},{"Lima","Sao Paulo"}}))
}
