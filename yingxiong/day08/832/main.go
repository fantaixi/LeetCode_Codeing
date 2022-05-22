package main

func main() {

}
func flipAndInvertImage(image [][]int) [][]int {
	m,n := len(image),len(image[0])
	for i := 0; i < m; i++ {
		for j,k := 0,n-1; j<=k ; j,k = j+1,k-1 {
			image[i][k],image[i][j] = 1 ^ image[i][j],1 ^ image[i][k]
		}
	}
	return image
}