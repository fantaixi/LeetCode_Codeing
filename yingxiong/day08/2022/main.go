package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Printf("第%d题：\n",i+1)
	}
}
func construct2DArray(original []int, m int, n int) [][]int {
	k := len(original)
	if k != m*n {
		return nil
	}
	ans := make([][]int, 0, m)
	for i := 0; i < k; i += n {
		ans = append(ans, original[i:i+n])
	}
	return ans
}