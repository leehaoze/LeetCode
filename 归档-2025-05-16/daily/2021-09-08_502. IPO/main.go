package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	type pair struct {
		Profit  int
		Capital int
	}

	projects := make([]pair, len(profits), len(profits))
	for i := range projects {
		projects = append(projects, pair{
			Profit:  profits[i],
			Capital: capital[i],
		})
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Capital < projects[j].Capital
	})

	h := &hp{}
	for cur := 0; k > 0; k-- {
		for cur < len(projects) && projects[cur].Capital <= w {
			heap.Push(h, projects[cur].Profit)
			cur++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
	}
	return w

}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func main() {
	fmt.Println(findMaximizedCapital(2, 0, []int{1,2,3},[]int{0,1,1}))
}
