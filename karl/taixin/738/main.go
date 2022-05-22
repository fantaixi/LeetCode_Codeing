package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(monotoneIncreasingDigits(332))
}
func monotoneIncreasingDigits(n int) int {
	if n < 10 {
		return n
	}
	num := []byte(strconv.Itoa(n))
	index,flag := check(num)
	if flag {
		return n
	}
	num[index]--
	for i := index+1; i < len(num); i++ {
		num[i] = '9'
	}
	res,_ := strconv.Atoi(string(num))
	return res
}
func check(num []byte) (int,bool){
	index,flag := 0,true
	for i := 1; i < len(num); i++ {
		if num[i] < num[index] {
			flag = false
			break
		}else if num[i] > num[index] {
			index = i
		}
	}
	return index,flag
}