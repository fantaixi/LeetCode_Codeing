package main

func main() {

}
func getConcatenation(nums []int) []int {
	newArr := make([]int,2*len(nums))
	for i := 0; i < len(nums); i++ {
		newArr[i] = nums[i]
		newArr[i+len(nums)] = nums[i]
	}
	return newArr
}
func getConcatenation1(nums []int) []int {
	return append(nums,nums...)
}