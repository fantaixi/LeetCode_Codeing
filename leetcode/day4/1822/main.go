package main

import "fmt"

func main() {
	arr := []int{-1,1,-1,1,-1}
	fmt.Println(arraySign(arr))
}
func arraySign(nums []int) int {
	tag := 0
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			tag++
		}
	}
	if tag % 2 == 0 {
		return 1
	}else {
		return -1
	}
}