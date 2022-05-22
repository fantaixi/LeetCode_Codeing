package main

import "math"

func main() {

}
func countKDifference(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if (float64(math.Abs(float64(nums[i]-nums[j]))) == float64(k)) {
				count ++
			}
		}
	}
	return count
}