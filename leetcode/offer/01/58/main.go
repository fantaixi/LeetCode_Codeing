package main

func main() {

}
func reverseLeftWords1(s string, n int) string {
	return s[n:]+s[:n]
}

func reverseLeftWords(s string, n int) string {
	for i := 0; i < len(s); i++ {
		s += string(s[i])
	}
	s = s[n:]
	return s
}

func reverseLeftWords3(s string, n int) string {
	b := []byte(s)
	// 1. 反转前n个字符
	// 2. 反转第n到end字符
	// 3. 反转整个字符
	reverse(b,0,n-1)
	reverse(b,n,len(b)-1)
	reverse(b,0,len(b)-1)
	return string(b)
}
func reverse(s []byte,left,right int) {
	for left < right {
		s[left],s[right] = s[right],s[left]
		left++
		right--
	}
}