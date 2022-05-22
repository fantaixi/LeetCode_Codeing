package main

func main() {

}
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	a,b,c := 0,1,1
	for i := 3; i <= n; i++ {
		a,b,c = b,c,a+b+c
	}
	return c
}