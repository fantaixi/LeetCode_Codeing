package main

import "fmt"

func main() {
	num := []int{1,2,3,4,5}
	for i := range num {
		fmt.Println(i)
	}
}
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	left,right := 1,1
	for right < length {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}