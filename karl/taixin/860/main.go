package main

import "fmt"

func main() {
	arr := []int{5,5,10,10,20}
	fmt.Println(lemonadeChange(arr))
}
func lemonadeChange(bills []int) bool {
	five,ten := 0,0
	for _,i := range bills {
		if i == 5 {
			five++
		}else if i == 10 {
			if five <= 0 {
				return false
			}
			five--
			ten++
		}else {
			if five > 0 && ten > 0 {
				five--
				ten--
			}else if five >= 3 {
				five -= 3
			}else {
				return false
			}
		}
	}
	return true
}