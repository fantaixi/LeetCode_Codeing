package _51

func reverseWords1(s string) string {
	str := []byte(s)
	length := len(str)
	slow, fast := 0, 0
	//去掉左边多余的空格
	for length > 0 && fast < length && str[fast] == ' ' {
		fast++
	}
	//去电中间多余的空格
	for ; fast < length; fast++ {
		if fast > 1 && str[fast] == str[fast-1] && str[fast] == ' ' {
			continue
		}
		str[slow] = str[fast]
		slow++
	}
	//去掉右边多余的空格
	if slow > 1 && str[slow-1] == ' ' {
		str = str[:slow-1]
	} else {
		str = str[:slow]
	}
	//反转整个字符串
	reverse1(str)
	i := 0
	for i < len(str) {
		//反转单个单词
		j := i
		//找到单词的结束位置
		for j < len(str) && str[j] != ' ' {
			j++
		}
		//反转
		reverse1(str[i:j])
		//更新下一个单词的起始位置，+1是要跳过单词间的空格
		i = j + 1
	}
	return string(str)
}
func reverse1(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
