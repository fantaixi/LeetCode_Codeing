package main

func main() {

}
func searchRange(nums []int, target int) []int {
	leftBorder := getLeft(nums,target)
	rightBorder := getRight(nums,target)
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1,-1}
	}
	if rightBorder-leftBorder > 1 {
		return []int{leftBorder+1,rightBorder-1}
	}
	return []int{-1,-1}
}

// 二分查找，寻找target的左边界leftBorder（不包括target）
// 如果leftBorder没有被赋值（即target在数组范围的右边，例如数组[3,3],target为4），为了处理情况一
func getLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	leftBorder := -2  // 记录一下leftBorder没有被赋值的情况
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {  // 寻找左边界，就要在nums[middle] == target的时候更新right
			right = mid - 1
			leftBorder = right
		} else {
			left = mid + 1
		}
	}
	return leftBorder
}

// 二分查找，寻找target的右边界（不包括target）
// 如果rightBorder为没有被赋值（即target在数组范围的左边，例如数组[3,3]，target为2），为了处理情况一
func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	rightBorder := -2
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1  // target 在左区间，所以[left, middle - 1]
		} else {  // 当nums[middle] == target的时候，更新left，这样才能得到target的右边界
			left = mid + 1
			rightBorder = left
		}
	}
	return rightBorder
}
