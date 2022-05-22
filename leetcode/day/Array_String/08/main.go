package main

func main() {

}
func setZeroes(matrix [][]int)  {
	row := make(map[int]bool)
	col := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] == true || col[j] == true{
				matrix[i][j] = 0
			}
		}
	}
}