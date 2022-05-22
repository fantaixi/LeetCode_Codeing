package main

func main() {

}
func maxProfit(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		result += max(prices[i]-prices[i-1],0)
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}