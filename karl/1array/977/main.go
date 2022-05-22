package main

func main() {

}
func sortedSquares(nums []int) []int {
	length := len(nums)
	res := make([]int,length)
	left,right,tips := 0,length-1,length-1
	//从后往前遍历新的数组
	for left <= right {
		leftV := nums[left]*nums[left]
		rightV := nums[right]*nums[right]
		if leftV > rightV {
			res[tips] = leftV
			left++
		}else {
			res[tips] = rightV
			right--
		}
		tips--
	}
	return res
}