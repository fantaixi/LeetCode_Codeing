package main

func main() {

}
func minOperations(nums []int) int {
	count:=0
	pre := nums[0]
	for i := 1; i < len(nums);i++ {
		if pre < nums[i] {
			pre = nums[i]
		}else {
			count += pre + 1 - nums[i]
			pre += 1
		}
	}
	return count
}