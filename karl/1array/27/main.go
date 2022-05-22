package main

func main() {

}

func removeElement(nums []int, val int) int {
	length := len(nums)
	res := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
//双指针优化
func removeElement1(nums []int, val int) int {
	left,right := 0,len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		}else {
			left++
		}
	}
	return left
}