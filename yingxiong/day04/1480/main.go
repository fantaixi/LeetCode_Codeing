package main

func main() {

}
func runningSum(nums []int) []int {
	temp := 0
	for i := 0; i < len(nums); i++ {
			temp += nums[i]
			nums[i] = temp
	}
	return nums
}