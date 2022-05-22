package main

func main() {

}
func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[0] == target
	}
	left,right := 0,n-1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] && nums[mid] ==nums[right]  {
			left++
			right--
		}else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid]  {
				right = mid - 1
			}else {
				left = mid + 1
			}
		}else {
			if nums[mid] < target && target <= nums[n-1] {
				left = mid+1
			}else {
				right = mid -1
			}
		}
	}
	return false
}