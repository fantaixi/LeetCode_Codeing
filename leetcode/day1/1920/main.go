package main

func main() {

}
func buildArray(nums []int) []int {
	newArr := make([]int,len(nums))
	for i := 0; i < len(nums); i++ {
		newArr[i] = nums[nums[i]]
	}
	return newArr
}