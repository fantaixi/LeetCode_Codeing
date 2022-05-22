package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1,4,3,2}
	fmt.Println(arrayPairSum(arr))
}
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sub := 0
	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 {
			sub += min(nums[i],nums[i+1])
		}
	}
	return sub
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}