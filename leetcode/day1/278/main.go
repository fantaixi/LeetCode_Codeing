package main

func main() {

}
func isBadVersion(version int) bool{
	return true
}
func firstBadVersion(n int) int {
	if n < 2 {
		return n
	}
	min,max := n,0
	for min != max+1 {
		mid := (min + max) / 2
		if isBadVersion(mid) {
			min = mid
		}else {
			max = mid
		}
	}
	return min
}
