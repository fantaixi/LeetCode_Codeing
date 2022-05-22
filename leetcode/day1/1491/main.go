package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{48000,59000,99000,13000,78000,45000,31000,17000,39000,37000,93000,77000,33000,28000,4000,54000,67000,6000,1000,11000}
	fmt.Println(average(arr))
	fmt.Println(average1(arr))
}
func average(salary []int) float64 {
	var sum int
	min := salary[0]
	max := -1
	for _, i := range salary {
		sum += i
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	total  := float64(sum - min - max)
	return total / float64(len(salary)-2)
}

func average1(salary []int) float64 {
	length := len(salary)
	sort.Ints(salary)
	min := salary[0]
	max := salary[length-1]
	sum :=0
	for i := 0; i < length; i++ {
		sum += salary[i]
	}
	return float64(1.0 * ((sum - min - max) / (length - 2)))
}