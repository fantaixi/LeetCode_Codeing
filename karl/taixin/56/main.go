package main

import (
	"fmt"
	"sort"
)

func main() {
	arr :=  [][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Println(merge(arr))
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	newArr := [][]int{}
	prev := intervals[0]
	//fmt.Println(prev)
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] {
			newArr = append(newArr,prev)
			prev = cur
		}else {
			prev[1] = max(prev[1],cur[1])
		}
	}
	newArr = append(newArr,prev)
	return newArr
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}