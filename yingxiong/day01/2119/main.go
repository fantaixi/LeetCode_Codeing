package main

import "fmt"

func main() {
	fmt.Println(1800%10)
}
func isSameAfterReversals(num int) bool {
	return num == 0 || num %10 > 0
}