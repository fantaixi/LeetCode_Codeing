package main

func main() {

}
func search1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left,right := 0,n-1
	for left <= right {
		mid := (left+right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			}else {
				left = mid + 1
			}
		}else {
			if nums[mid] < target && target <= nums[n-1] {
				left = mid + 1
			}else {
				right = mid - 1
			}
		}
	}
	return -1
}
