package _91
/*
给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
示例:
输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
说明:
给定数组的长度不会超过15。
数组中的整数范围是 [-100,100]。
给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
 */
/*
提示：
1 <= nums.length <= 15
-100 <= nums[i] <= 100
 */
func findSubsequences(nums []int) [][]int {
	var path []int
	var res [][]int
	backtrack(0,nums,path,&res)
	return res
}
func backtrack(startIndex int, nums, path []int, res *[][]int) {
	if len(path) > 1 {
		temp := make([]int,len(path))
		copy(temp,path)
		*res = append(*res,temp)
	}
	//记录本层元素使用记录
	history := [201]int{}
	for i := startIndex; i < len(nums); i++ {
		//分两种情况判断：一，当前取的元素小于子集的最后一个元素，则继续寻找下一个适合的元素
		//                或者二，当前取的元素在本层已经出现过了，所以跳过该元素，继续寻找
		if len(path) > 0 && nums[i] < path[len(path)-1] || history[nums[i]+100] == 1 {
			continue
		}
		//表示本层该元素使用过了
		history[nums[i]+100]=1
		path = append(path,nums[i])
		backtrack(i+1,nums,path,res)
		path = path[:len(path)-1]
	}
}
