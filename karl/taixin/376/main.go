package main

func main() {

}
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	result := 1
	preDiff := nums[1]-nums[0]
	if preDiff != 0 {
		result = 2
	}
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && preDiff <= 0 || diff < 0 && preDiff >= 0 {
			result++
			preDiff = diff
		}
	}
	return result
}
