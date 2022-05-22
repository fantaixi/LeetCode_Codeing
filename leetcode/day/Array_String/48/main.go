package main

func main() {

}
func rotate(matrix [][]int)  {
	n := len(matrix)
	temp := make([][]int,n)
	for i := range temp {
		temp[i] = make([]int,n)
	}
	for i, row := range matrix {
		for j, v := range row {
			temp[j][n-1-i] = v
		}
	}
	copy(matrix,temp)
}