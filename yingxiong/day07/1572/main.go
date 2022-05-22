package main

func main() {

}
func diagonalSum1(mat [][]int) int {
	n,count,mid :=len(mat),0, len(mat)/2
	for i := 0; i < n; i++ {
		count += mat[i][i] + mat[i][n-1-i]
	}
	return count - mat[mid][mid] * (n&1)
}

func diagonalSum(mat [][]int) int {
	sum,n := 0,len(mat)
	for i := 0; i < n; i++ {
		sum += mat[i][i]
	}
	for i,j :=n-1, 0; j<n ; i ,j= i-1,j+1 {
		sum += mat[i][j]
	}
	if n%2 == 0 {
		return sum
	}
	return sum -mat[n/2][n/2]
}