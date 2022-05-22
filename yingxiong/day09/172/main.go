package main

func main() {

}
func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}
	return n/5 + trailingZeroes(n/5)
}