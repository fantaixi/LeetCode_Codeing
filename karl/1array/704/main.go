package main

import "fmt"

func main() {
	arr := []int{-1,0,3,5,9,12}
	fmt.Println(search1(arr,9))
}
func search(nums []int, target int) int {
	left,right := 0,len(nums)
	for left < right {
		mid := left + (right - left) >> 1
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			right = mid
		}else {
			left  = mid + 1
		}
	}
	return -1
}

func search1(nums []int, target int) int {
	left,right := 0,len(nums)-1
	for left <= right {
		mid := left + (right - left) >> 1
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	return -1
}