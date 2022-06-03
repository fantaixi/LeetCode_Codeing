package _7
/*
编写一个程序，通过填充空格来解决数独问题。
一个数独的解法需遵循如下规则： 数字 1-9 在每一行只能出现一次。 数
字 1-9 在每一列只能出现一次。 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。 空白格用 '.' 表示。
提示：
给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。
 */
func solveSudoku(board [][]byte)  {
	var dfs func(board [][]byte)bool
	dfs = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				//判断此位置是否适合填数字
				if board[i][j] != '.' {
					continue
				}
				//尝试填1-9
				for k := '1'; k <= '9'; k++ {
					if isValid(i, j, byte(k), board) == true { //如果满足要求就填
						board[i][j] = byte(k)
						if dfs(board) == true {
							return true
						}
						board[i][j]='.'
					}
				}
				return false
			}
		}
		return true
	}
	dfs(board)
}
//判断填入数字是否满足要求
func isValid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ { //行
		if board[row][i] == k {
			return false
		}
	}
	for i := 0; i < 9; i++ { //列
		if board[i][col] == k {
			return false
		}
	}
	//方格
	startRow := (row/3)*3
	startCol := (col/3)*3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}