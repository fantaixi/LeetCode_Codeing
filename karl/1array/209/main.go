package main

import "math"

func main() {

}
//滑动窗口
func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	start,end := 0,0
	ans,sum := math.MaxInt32,0
	for end < length {
		sum += nums[end]
		for sum >= target {
			//求最短
			ans = min(ans,end-start+1)
			//继续移动start
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}