package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int)
	for idx, el := range list2 {
		m[el] = idx
	}

	ret := make([]string, 0)
	minIdx := len(list1)*2 + 1

	for idx, el := range list1 {
		if idx2, exists := m[el]; exists {
			if minIdx == idx+idx2 {
				ret = append(ret, el)
			} else if minIdx > idx+idx2 {
				ret = make([]string, 0)
				ret = append(ret, el)
				minIdx = idx + idx2
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
}
