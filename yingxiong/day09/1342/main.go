package main

func main() {

}
func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	if num %2 == 1 {
		return numberOfSteps(num-1)+1
	}else {
		return numberOfSteps(num/2)+1
	}
}