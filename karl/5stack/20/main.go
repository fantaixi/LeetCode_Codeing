package _0
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
 */
func isValid(s string) bool {
	stack := []string{}
	for _, ch := range s {
		c := string(ch)
		if c == "{" || c == "(" || c == "[" {
			stack = append(stack,c)
		}else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top == "(" && c == ")" || top == "[" && c == "]" || top == "{" && c == "}" {
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		}
	}
	return len(stack)==0
}
