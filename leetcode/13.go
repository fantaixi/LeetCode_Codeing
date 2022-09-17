package main

/*
罗马数字转整数
*/
func romanToInt(s string) int {
	/// 用map做映射
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	length := len(s)
	ans := m[s[length-1]]
	// 从右到左遍历每个字符，检查当前字符的映射值，是否大于上一个，是则加当前的值，否则，减去当前的值。
	for i := length - 2; i >= 0; i-- {
		if m[s[i]] >= m[s[i+1]] {
			ans += m[s[i]]
		} else {
			ans -= m[s[i]]
		}
	}
	return ans
}
