package main

import "sort"

func main() {

}
func wiggleSort(nums []int)  {

	l := len(nums)
	left, right := l, (l+1)>>1
	t := make([]int, l)
	copy(t, nums)
	sort.Ints(t)
	for i := 0; i < l; i++ {
		if i&1 == 1 {
			left--
			nums[i] = t[left]
		} else {
			right--
			nums[i] = t[right]
		}
	}
}