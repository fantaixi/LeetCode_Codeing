package main

func main() {

}
func transpose(matrix [][]int) [][]int {
	lr := len(matrix)
	lc := len(matrix[0])
	result := make([][]int,lc)
	for i, _ := range result {
		result[i] = make([]int,lr)
	}
	for i := 0; i < lr; i++ {
		for j := 0; j < lc; j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}