package main

func main() {

}
func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	k := 1
	for i := 1; i <= 15; i++ {
		k *= 4
		if k == n {
			return true
		}
	}
	return false
}