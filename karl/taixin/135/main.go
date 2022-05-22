package main

func main() {

}
func candy(ratings []int) int {
	arr := make([]int,len(ratings))
	//确保每个孩子都有糖果
	for i := 0; i < len(arr); i++ {
		arr[i] = 1
	}
	//从左到右，当右边的大就+1
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			arr[i+1] = arr[i]+1
		}
	}
	//从右到左，左边的大就+1
	for i := len(ratings)-1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			arr[i-1] = max(arr[i-1],arr[i]+1)
		}
	}
	//计算总数
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}