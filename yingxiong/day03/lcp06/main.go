package main

func main() {

}
func minCount(coins []int) int {
	count := 0
	for i := 0; i < len(coins); i++ {
			count +=  coins[i] / 2 + coins[i] %2
	}
	return count
}