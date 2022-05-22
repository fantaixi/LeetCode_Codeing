package main

import "sort"

func main() {

}
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	child := 0
	for sIndex := 0; child < len(g) && sIndex < len(s); sIndex++ {
		if s[sIndex] >= g[child] {
			child++
		}
	}
	return child
}