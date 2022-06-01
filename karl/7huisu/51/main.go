package _1

import "strings"

/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

皇后们的约束条件：
不能同行
不能同列
不能同斜线
*/
/*
棋盘的宽度就是for循环的长度，递归的深度就是棋盘的高度
 */
func solveNQueens(n int) [][]string {
	res := [][]string{}
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	var dfs func(int)
	dfs = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i, rowStr := range chessboard {
				temp[i] = strings.Join(rowStr, "")
			}
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			if isValid(n, row, i, chessboard) {
				chessboard[row][i] = "Q" // 放置皇后
				dfs(row + 1)
				chessboard[row][i] = "." // 回溯，撤销皇后
			}
		}
	}
	dfs(0)
	return res
}
// n 为输入的棋盘大小
// row 是当前递归到棋盘的第几行了
func isValid(n, row, col int, chessboard [][]string) bool {
	// 检查列
	for i := 0; i < row; i++ {  // 这是一个剪枝
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	// 检查 45度角是否有皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	// 检查 135度角是否有皇后
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}
