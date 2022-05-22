package main

import (
	"fmt"
	"sort"
)

func main() {
	for i := 0; i < 9; i++ {
		fmt.Printf("第%d题：\n",i+1)
	}
}
func numRescueBoats(people []int, limit int) int {
	count := 0
	sort.Ints(people)
	left,right := 0,len(people)-1
	for left <= right {
		if people[left]+people[right] > limit {
			right --
		}else {
			left++
			right--
		}
		count++
	}
	return count
}