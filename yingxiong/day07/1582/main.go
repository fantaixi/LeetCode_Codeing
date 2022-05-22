package main

func main() {

}
func numSpecial(mat [][]int) int {
	m,n,num := len(mat),len(mat[0]),0
	row := make([]int,m)
	col := make([]int,n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] ==1  && row[i] == 1 && col[j] ==1{
				num++
			}
		}
	}
	return num
}
