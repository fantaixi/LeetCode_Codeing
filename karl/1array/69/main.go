package main


func main() {

}
func mySqrt(x int) int {
	left,right := 0,x
	res := -1
	for left <= right {
		mid := left + (right - left) >> 1
		if mid * mid <= x {
			res = mid
			left = mid + 1
		}else {
			right = mid - 1
		}
	}
	return res
}
