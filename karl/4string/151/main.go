package _51
/*
给你一个字符串 s ，颠倒字符串中 单词 的顺序。
单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
 */

/*
移除多余空格
将整个字符串反转
将每个单词反转
 */
func reverseWords(s string) string {
	str := []byte(s)
	slowIndex,fastIndex := 0,0
	for len(str) > 0 && fastIndex < len(str) && str[fastIndex] == ' ' {
		fastIndex++
	}
	for ; fastIndex < len(str); fastIndex++ {
		if fastIndex-1>0 && str[fastIndex-1] == str[fastIndex] && str[fastIndex]==' '{
			continue
		}
		str[slowIndex] = str[fastIndex]
		slowIndex++
	}
	if slowIndex-1>0 && str[slowIndex-1] ==' ' {
		str=str[:slowIndex-1]
	}else {
		str = str[:slowIndex]
	}
	reverse(&str,0,len(str)-1)
	i := 0
	for i < len(str) {
		j:=i
		for ; j < len(str) && str[j] != ' '; j++ {
			reverse(&str,i,j-1)
			i=j
			i++
		}
	}
	return string(str)
}
func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left],(*b)[right] = (*b)[right],(*b)[left]
		left++
		right--
	}
}
