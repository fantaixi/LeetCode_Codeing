package main

func main() {

}
func isPowerOfTwo1(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	k := 1
	for i := 1; i <= 31 ; i++ {
		k *= 2
		if k == n {
			return true
		}
	}
	return false
}

func isPowerOfTwo(n int) bool {
	const big = 1 << 30
	return n >0 && big % n ==0
}
