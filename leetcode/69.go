package main

/*
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
*/
func mySqrt(x int) int {
	left, right := 1, x/2+1
	for left <= right {
		mid := left + (right-left)>>1
		if mid == x/mid {
			return mid
		} else if mid > x/mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
