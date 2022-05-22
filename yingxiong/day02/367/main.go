package main

func main() {

}
func isPerfectSquare1(num int) bool {
	sub := 0
	for i := 1; ; i++ {
		sub  = i * i
		if sub == num {
			return true
		}
		if sub > num {
			return false
		}
	}
	return false
}

func isPerfectSquare2(num int) bool {
	for i := 1; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
func isPerfectSquare(num int) bool {
	left,rigrht := 0,num
	for left <= rigrht {
		mid := (left +rigrht) >> 1
		sub := mid * mid
		if sub < num {
			left = mid +1
		}else if sub > num {
			rigrht = mid -1
		}else {
			return true
		}
	}
	return false
}
