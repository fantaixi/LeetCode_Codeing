package main

import "sort"

func main() {

}
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	cont := make(map[int]int)
	for _, val := range nums {
		if _, ok := cont[val]; ok {
			return true
		}else {
			cont[val] = 1
		}
	}
	return false
}