package _44

/*
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
 */
func reverseString(s []byte)  {
	for left,right := 0,len(s)-1; left<right; left,right = left+1,right-1 {
		s[left],s[right] = s[right],s[left]
	}
}

func reverseString1(s []byte)  {
	left,right := 0,len(s)-1
	for left < right {
		s[left],s[right]=s[right],s[left]
		left++
		right--
	}
}

