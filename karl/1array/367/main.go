package main

func main() {

}
func isPerfectSquare(num int) bool {
	left,right := 0,num
	for left <= right {
		mid := left + (right - left) >> 1
		if mid*mid < num {
			left = mid + 1
		}else if mid*mid > num {
			right = mid-1
		}else{
			return true
		}
	}
	return false
}