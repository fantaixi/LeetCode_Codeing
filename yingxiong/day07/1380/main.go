package main

func main() {

}
func luckyNumbers (matrix [][]int) []int {
	minRow := make([]int,len(matrix))
	maxCol := make([]int,len(matrix[0]))
	for i, row := range matrix {
		minRow[i] = row[0]
		for j, x := range row {
			minRow[i] = min(minRow[i],x)
			maxCol[j] = max(maxCol[j],x)
		}
	}
	for i, row := range matrix {
		for j, x := range row {
			if x == minRow[i] && x == maxCol[j] {
				return []int{x}
			}
		}
	}
	return nil
}
func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}