package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(subtractProductAndSum(3210))
}
func subtractProductAndSum(n int) int {
	//先把数字转换成字符串再截取
	data := strconv.Itoa(n)
	sum := 0
	mult := 1
	for i := 0; i < len(data); i++ {
		s , _ :=strconv.Atoi(data[i:i+1])
		//fmt.Println(s)
		sum += s
		mult *= s

	}
	fmt.Println(mult)
	fmt.Println(sum)
	return mult -sum
}
