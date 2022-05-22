package main

func main() {

}
func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i, v := range nums {
		if 2*sum + v == total {
			return i
		}
		sum += v
	}
	return -1
}