package offer58
/*
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
 */
func reverseLeftWords(s string, n int) string {
	for i := 0; i < n; i++ {
		s += string(s[i])
	}
	s = s[n:]
	return s
}

func reverseLeftWords1(s string, n int) string {
	str := []byte(s)
	reverse(str,0,n-1)
	reverse(str,n,len(str)-1)
	reverse(str,0,len(str)-1)
	return string(str)
}
func reverse(s []byte, left, right int) {
	for left < right {
		s[left],s[right] = s[right],s[left]
		left++
		right--
	}
}
