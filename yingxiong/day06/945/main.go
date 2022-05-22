package main

import "sort"

func main() {

}
func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			count += nums[i] - nums[i+1] +1
			nums[i+1] += nums[i] - nums[i+1]+1
		}
	}
	return count
}