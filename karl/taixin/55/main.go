package main

func main() {

}
func canJump(nums []int) bool {
	cover := 0
	if len(nums) == 0 {
		return true
	}
	//看i的活动范围，取里面的最大值，能覆盖到最后一个元素说明就能成功
	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i],cover)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}