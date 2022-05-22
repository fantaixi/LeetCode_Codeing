package main

func main() {

}
func countNegatives1(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] <0 {
				count++
			}
		}
	}
	return count
}

func countNegatives(grid [][]int) int {
	count := 0
	for _, val := range grid {
		count += binary(val)
	}
	return count
}
func binary(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= 0 {
			left++
		} else {
			if mid == 0 || nums[mid-1] >= 0 {
				return len(nums) - mid
			}
			right--
		}
	}
	return 0
}
