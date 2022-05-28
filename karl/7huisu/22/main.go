package _2
/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 */
func generateParenthesis(n int) []string {
	res = []string{}
	backtrack(n,0,0,"")
	return res
}
var res []string
func backtrack(n,left,right int, s string) {
	//如果有效数量等于n*2，直接返回
	if len(s) == n*2 {
		res = append(res,s)
		return
	}else {
		//如果左括号小于n的话，试探添加左括号
		if left < n {
			backtrack(n,left+1,right,s+"(")
		}
		//如果左括号的数量大于右括号，试探添加右括号
		if left > right {
			backtrack(n,left,right+1,s+")")
		}
	}
}