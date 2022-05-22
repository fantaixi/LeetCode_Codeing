package main

func main() {

}
func game(guess []int, answer []int) int {
	count := 0
	for i := 0; i < 3; i++ {
			if guess[i] == answer[i] {
				count++
			}
	}
	return count
}