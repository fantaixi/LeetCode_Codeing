package main

func main() {

}
func partitionLabels(s string) []int {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}
	res := []int{}
	left,right := 0,0
	for i := 0; i < len(s); i++ {
		tempPos := m[s[i]]
		if tempPos > right {
			right = tempPos
		}
		if i == right {
			res = append(res,i-left+1)
			left = i + 1
		}
	}
	return res
}