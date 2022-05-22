package main

func main() {

}
func reverseLeftWords1(s string, n int) string {
	return s[n:]+s[0:n]
}
func reverseLeftWords(s string, n int) string {
	for i := 0; i < n; i++ {
		s += string(s[i])
	}
	s = s[n:]
	return s
}