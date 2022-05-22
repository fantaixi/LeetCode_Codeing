package main

func main() {

}
func createTargetArray(nums []int, index []int) []int {
	target := make([]int,len(nums))
	for m, n := range index {
		copy(target[n+1:],target[n:])
		target[n] = nums[m]
	}
	return target
}