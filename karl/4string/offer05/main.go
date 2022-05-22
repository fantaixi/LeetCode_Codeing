package offer05
/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
 */
//遍历添加
func replaceSpace(s string) string {
	newStr := ""
	for _, val := range s {
		if val == ' ' {
			newStr = "%20"
		}else {
			newStr += string(val)
		}
	}
	return newStr
}
//原地修改
func replaceSpace1(s string) string {
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	for _, val := range b {
		if val == ' ' {
			spaceCount++
		}
	}
	temp := make([]byte,2*spaceCount)
	b = append(b,temp...)
	i,j := length-1,len(b)-1
	for i >= 0 {
		if b[i] != ' ' {
			b[j]=b[i]
			i--
			j--
		}else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j -= 3
		}
	}
	return string(b)
}
