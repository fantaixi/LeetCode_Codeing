package main

import "fmt"

/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
 */

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n-1; i >= 0; i-- {
		//如果不等于9就直接+1
		if digits[i] != 9 {
			digits[i] ++
			//等于9就让后面一位变为0
			for j := i+1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	//全部为9就直接扩大一位
	digits = make([]int,n+1)
	digits[0] = 1
	return digits
}

/*
换个思路
从后往前遍历，每位+1，如果+1之后是10的话，肯定变成0，变成0就要一直往前进位，迭代不能停，一直迭代下去直到不为10，
如果不能当场返回，那肯定是最高位也变成0，最后直接在最前面加个1
 */
func plusOne1(digits []int) []int {
	for i := len(digits)-1; i >= 0 ; i-- {
		digits[i] += 1
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	return append([]int{1},digits...)
}
func main() {
	arr := []int{1,2,9}
	fmt.Println(plusOne(arr))

}
