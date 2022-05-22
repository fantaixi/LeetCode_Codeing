package main

import "sort"

func main() {

}
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	sub := 0
	for i := 0; i < len(nums)-1; i++ {
		v := nums[i+1] - nums[i]
		if v > sub {
			sub = v
		}
	}
	return sub
}