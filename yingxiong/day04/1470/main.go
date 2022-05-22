package main

func main() {

}
func shuffle(nums []int, n int) []int {
	index := 0
	newArr := make([]int,2*n)
	for i := 0; i < n; i++ {
			newArr[i*2] = nums[index]
			newArr[i*2+1] = nums[index+n]
			index++
	}
	return newArr
}
