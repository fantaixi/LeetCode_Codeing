package main

import "fmt"

func main() {
	fmt.Println(sumNums(6))
}
func sumNums1(n int) int {
	if n == 1 {
		return 1
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}
func sumNums2(n int) int {
	_ = n > 0 && func() bool { n += sumNums2(n - 1); return true }()
	return n
}
