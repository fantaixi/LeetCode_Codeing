package _50

import "strconv"

/*
根据 逆波兰表示法，求表达式的值。
有效的算符包括+、-、*、/。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
注意两个整数之间的除法只保留整数部分。
可以保证给定的逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
 */
/*
如果遇到操作数，则将操作数入栈；
如果遇到运算符，则将两个操作数出栈，
其中先出栈的是右操作数，后出栈的是左操作数，使用运算符对两个操作数进行运算，将运算得到的新操作数入栈。

 */
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val,err := strconv.Atoi(token)
		if err==nil {
			stack = append(stack,val)
		}else {
			n1,n2 := stack[len(stack)-2],stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack,n1+n2)
			case "-":
				stack = append(stack,n1-n2)
			case "*":
				stack = append(stack,n1*n2)
			case "/":
				stack = append(stack,n1/n2)
			}
		}
	}
	return stack[0]
}
