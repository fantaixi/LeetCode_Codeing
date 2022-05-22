package main

import "fmt"

func main() {
	nums := []int{1,3}
	fmt.Println(searchInsert1(nums,7))
}
//使用二分
//方法1： target 是在一个在左闭右闭的区间里，也就是[left, right]
func searchInsert(nums []int, target int) int {
	left,right := 0,len(nums)-1
	for left <= right {  //当left==right，区间[left, right]依然有效
		mid := left + (right - left) >> 1  //防止溢出 等同于(left + right)/2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			right = mid - 1  // target 在左区间，所以[left, middle - 1]
		}else {
			left = mid + 1  // target 在右区间，所以[middle + 1, right]
		}
	}
	return right+1
}
//方法2： target 是在一个在左闭右开的区间里，也就是[left, right)
func searchInsert1(nums []int, target int) int {
	left,right := 0,len(nums)
	for left < right {  //当left==right，区间[left, right]无效
		mid := left + (right - left) >> 1
		num := nums[mid]
		if num == target {
			return mid
		}else if num > target {
			right = mid   // target 在左区间，在[left, middle)中
		}else {
			left = mid + 1  // target 在右区间，在 [middle+1, right)中
		}
	}
	return right
}
//暴力
func searchInsert2(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}