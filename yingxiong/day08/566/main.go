package main

func main() {

}
func matrixReshape(mat [][]int, r int, c int) [][]int {
	 m,n := len(mat),len(mat[0])
	if m*n != r*c {
		return mat
	}
	res := make([][]int,r)
	for i := range res {
		res[i] = make([]int,c)
	}
	for i := 0; i < n*m; i++ {
		res[i/c][i%c] = mat[i/n][i%n]
	}
	return res
}