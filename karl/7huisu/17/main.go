package _7
/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 */
func letterCombinations(digits string) []string {
	length := len(digits)
	res := []string{}
	if length < 1 {
		return res
	}
	digitsMap := [10]string{
		"", // 0
		"", // 1
		"abc", // 2
		"def", // 3
		"ghi", // 4
		"jkl", // 5
		"mno", // 6
		"pqrs", // 7
		"tuv", // 8
		"wxyz", // 9
	}
	backstrack("",digits,0,digitsMap,&res)
	return res
}
func backstrack(tempString, digits string, index int, digitsMap [10]string, res *[]string) {
	//终止条件，字符串长度等于digits的长度
	if len(tempString) == len(digits) {
		*res = append(*res,tempString)
		return
	}
	tempK := digits[index]-'0'  // 将index指向的数字转为int（确定下一个数字）
	letter := digitsMap[tempK]  // 取数字对应的字符集
	for i := 0; i < len(letter); i++ {
		tempString = tempString+string(letter[i])  //拼接结果
		backstrack(tempString,digits,index+1,digitsMap,res)
		tempString=tempString[:len(tempString)-1]  //回溯
	}
}
