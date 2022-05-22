package main

func main() {

}
func searchRange1(nums []int, target int) []int {
	// 目标值开始位置：为 target 的左侧边界
	start := leftBorder(nums,target)
	// 如果开始位置越界 或 目标值不存在，直接返回
	if start == len(nums) || nums[start] != target {
		return []int{-1,-1}
	}
	// 目标值结束位置：为 target 的右侧边界
	end := rightBorder(nums,target)
	return []int{start,end}
}
func leftBorder(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			right = mid - 1 //收紧右侧边界以锁定左侧边界
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {  //因为有target相同的情况，所以这个elseif不能省
			right = mid - 1
		}
	}
	//返回左侧边界
	return left
}
func rightBorder(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			left = mid + 1 //收紧左侧边界以锁定右侧边界
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {  //因为有target相同的情况，所以这个elseif不能省
			right = mid - 1
		}
	}
	//返回右侧边界
	return right
}
