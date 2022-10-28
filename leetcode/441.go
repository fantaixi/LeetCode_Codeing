package main

/*
你总共有n枚硬币，并计划将它们按阶梯状排列。对于一个由 k 行组成的阶梯，其第 i 行必须正好有 i 枚硬币。阶梯的最后一行 可能 是不完整的。

给你一个数字n ，计算并返回可形成 完整阶梯行 的总行数。
 */
func arrangeCoins(n int) int {
	left,right := 1,n
	for left <= right {
		mid := left + (right-left) >> 1
		if (mid+1)*mid/2 <= n {
			left = mid+1
		}else {
			right = mid-1
		}
	}
	return right
}
