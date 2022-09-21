package main

/*
给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。
 */
func search(nums []int, target int) int {
	left,right := 0,len(nums) -1
	for left <= right {
		mid := left + (right-left) >> 1
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			right = mid-1
		}else {
			left = mid+1
		}
	}
	return -1
}
