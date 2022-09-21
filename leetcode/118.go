package main

/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。
 */
func generate(numRows int) [][]int {
	ans := make([][]int,numRows)
	for i, _ := range ans {
		ans[i] = make([]int,i+1)
		ans[i][0] = 1
		ans[i][i] =1
		for j := 1; j < i; j++ {
			//每个数是它左上方和右上方的数的和
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}