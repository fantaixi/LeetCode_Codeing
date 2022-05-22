package main

func main() {

}
func maxSlidingWindow(nums []int, k int) []int {
	var stack []int
	var res []int
	for i, v := range nums {
		for len(stack) > 0 && v >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,i)
		if i-k+1 > stack[0] {
			stack = stack[1:]
		}
		if i+1 >= k {
			res = append(res,nums[stack[0]])
		}
	}
	return res
}
func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) <= 0 {
		return nil
	}
	res := make([]int,len(nums)-k+1)
	max := 0
	//先找出第一组中最小的
	for i := 0; i < k; i++ {
		if nums[max] < nums[i] {
			max = i
		}
	}
	res[0] = nums[max]
	//变量nums，依次找出每个组中最大的
	for i := k; i < len(nums); i++ {
		// 如果最大的下标没有在当前组中，遍历新数组找到最大的下标
		if max <= i-k {
			max = i-k+1
			for j := max+1; j <= i; j++ {
				if nums[max] < nums[j] {
					max = j
				}
			}
		}else if nums[max] < nums[i] { //// 最大的在新的组中，和新加入组中的比较，获得最大下标
			max = i
		}
		res[i-k+1] = nums[max]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}