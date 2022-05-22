package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4,2,3}
	fmt.Println(largestSumAfterKNegations(arr,1))
}
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k >0 {
			nums[i] *= -1
			k--
		}
	}
	sort.Ints(nums)
	if k%2 == 1 {
		nums[0] = -nums[0]
	}
	for _, v := range nums {
		sum += v
	}
	return sum
}