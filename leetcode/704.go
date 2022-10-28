package main

/*
给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。
*/
//左闭右闭
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1 //因为是左闭右闭，所以不能包含已经比较过的下标
		} else {
			left = mid + 1
		}
	}
	return -1
}

//左闭右开
func search1(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

//func main() {
//	nums := []int{-1,0,3,5,9,12}
//	fmt.Println(search1(nums,9))
//}
