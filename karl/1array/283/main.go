package main

func main() {

}
func moveZeroes(nums []int)  {
	left,right := 0,0
	for right < len(nums) {
		//right负责寻找非零数,找到就与左指针对应的值互换
		if nums[right] != 0 {
			nums[left],nums[right] = nums[right],nums[left]
			left++
		}
		right++
	}
}