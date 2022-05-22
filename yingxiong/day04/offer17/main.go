package main

func main() {

}
func printNumbers(n int) []int {
	temp := 1
	for i := 0; i < n; i++ {
		temp *= 10
	}
	arr := make([]int,temp-1)
	for i := 1; i < temp; i++ {
		arr[i-1] = i
	}
	return arr
}