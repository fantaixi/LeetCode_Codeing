package main

func main() {

}
func maxSubArray(nums []int) int {
	result,temp := 0,0
	result = nums[0]
	for i := 1; i <= len(nums); i++ {
		temp += nums[i-1]
		if temp > result {
			result = temp
		}
		if temp <= 0 {
			temp = 0
		}
	}
	return result
}