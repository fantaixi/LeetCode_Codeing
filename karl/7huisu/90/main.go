package _0

import "sort"

/*
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
 */
/*
对78的去重
 */
var res [][]int
func subsetsWithDup(nums []int) [][]int {
	res = [][]int{}
	//有乱序的数组，要先排序
	sort.Ints(nums)
	backtrack([]int{},nums,0)
	return res
}
func backtrack(path, nums []int, start int) {
	temp := make([]int,len(path))
	copy(temp,path)
	res = append(res,temp)
	for i := start; i < len(nums); i++ {
		//遇到相同的就跳过
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path,nums[i])
		backtrack(path,nums,i+1)
		path = path[:len(path)-1]
	}
}
