package main

import "sort"

func main() {

}
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	dc := arr[1] - arr[0]
	for i := len(arr)-1; i > 0 ; i-- {
		if arr[i] - arr[i-1] != dc {
			return false
		}
	}
	return true
}