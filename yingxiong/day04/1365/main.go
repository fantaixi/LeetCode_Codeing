package main

import "sort"

func main() {

}
func smallerNumbersThanCurrent1(nums []int) []int {
	temp,count := 0,0
	newArr := make([]int,len(nums))
	for i := 0; i < len(nums); i++ {
		temp = nums[i]
		for j := 0; j < len(nums); j++ {
			if temp > nums[j] {
				count++
			}
		}
		newArr[i] = count
		count=0
	}
	return newArr
}

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int,len(nums))
	newArr := make([]int,len(nums))
	copy(newArr,nums)
	newMap := make(map[int]int)
	sort.Ints(nums)
	for i := len(nums)-1; i >=0 ; i-- {
		newMap[nums[i]] = i
	}
	for i, val := range newArr {
		res[i] = newMap[val]
	}
	return res
}