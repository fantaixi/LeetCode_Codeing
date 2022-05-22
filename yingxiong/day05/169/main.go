package main

import "fmt"

func main() {
	arr := []int{2,2,1,1,1,2,2,3,3,3,3,3,3}
	fmt.Println(majorityElement(arr))
}
func majorityElement(nums []int) int {
	condidata,count := 0,0
	for _, val := range nums {
		if count == 0 {
			condidata = val
		}
		if val == condidata {
			count++
		}else {
			count--
		}
	}
	return condidata
}