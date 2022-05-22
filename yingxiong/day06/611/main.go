package main

import "sort"

func main() {

}
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	count := 0
	n := len(nums)
	for i, v := range nums {
		k := i
		for j :=i + 1; j < n; j++ {
			for k+1 < n && nums[k+1] < v + nums[j]{
				k++
			}
			count += max(k-j,0)
		}
	}
	return count
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}