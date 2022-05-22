package main

func main() {

}
func sortedSquares(nums []int) []int {
	newArr := make([]int,0)
	for i := 0; i < len(nums); i++ {
		newArr = append(newArr,nums[i]*nums[i])
	}
	for i := 0; i < len(newArr); i++ {
		for j := 0; j < len(newArr) - i -1; j++ {
			if newArr[j] > newArr[j+1] {
				newArr[j+1],newArr[j] = newArr[j],newArr[j+1]
			}
		}
	}
	return newArr
}

func sortedSquares1(nums []int) []int {
	//每次比较两个指针对应的数，选择较大的那个逆序放入答案并移动指针
	n := len(nums)
	newArr := make([]int,n)
	i,j := 0,n-1
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := nums[i]*nums[i], nums[j]*nums[j]; v > w {
			newArr[pos] = v
			i++
		}else {
			newArr[pos] = w
			j--
		}
	}
	return newArr
}
