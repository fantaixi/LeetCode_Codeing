package main

/*
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。
 */
func getRow(rowIndex int) []int {
	ans := make([][]int,rowIndex+1)
	for i, _ := range ans {
		ans[i] = make([]int,i+1)
		ans[i][0]=1
		ans[i][i]=1
		for j := 1; j < i; j++ {
			ans[i][j]=ans[i-1][j]+ans[i-1][j-1]
		}
	}
	return ans[rowIndex]
}

/*
对第 i+1 行的计算仅用到了第 i 行的数据，因此可以使用滚动数组的思想优化空间复杂度
 */
func getRow1(rowIndex int) []int {
	var pre,cur []int
	for i := 0; i <= rowIndex; i++ {
		cur = make([]int,i+1)
		cur[0],cur[i] = 1,1
		for j := 1; j < i; j++ {
			cur[j] = pre[j]+pre[j-1]
		}
		pre = cur
	}
	return pre
}
//func main() {
//	fmt.Println(getRow1(3))
//}
