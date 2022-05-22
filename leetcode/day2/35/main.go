package main

import "fmt"

func main() {
	arr := []int{1,3,5,6}
	fmt.Println(searchInsert(arr,20))
}
func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	left,right := 0,len(nums)-1
	for left <= right{
		mid := left + (right - left) >> 1
		if nums[mid] < target {
			left = mid +1
		}else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left
}


