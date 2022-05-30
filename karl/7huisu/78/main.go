package _8

import "sort"

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
 */
/*
遍历树的时候，把每个节点都记录下来，就是所求
 */
var res [][]int
func subsets(nums []int) [][]int {
	res = make([][]int,0)
	path := []int{}
	sort.Ints(nums)
	backtrack(path,nums,0)
	return res
}

func backtrack(path,nums []int,startIndex int)  {
	temp := make([]int,len(path))
	copy(temp,path)
	res = append(res,temp)
	for i := startIndex; i < len(nums); i++ {
		path = append(path,nums[i])
		backtrack(path,nums,i+1)
		path = path[:len(path)-1]
	}
}
/*
逐个考察数字，每个数都选或不选。等到递归结束时，把集合加入解集。
 */
func subsets1(nums []int) [][]int {
	res := [][]int{}
	var dfs func(start int,path []int)
	dfs = func(start int, path []int) {
		//指针越界
		if start == len(nums) {
			temp := make([]int,len(path))
			copy(temp,path)
			res = append(res,temp)
			return
		}
		// 选择这个数
		path = append(path,nums[start])
		// 基于该选择，继续往下递归，考察下一个数
		dfs(start+1,path)
		// 上面的递归结束，撤销该选择
		path = path[:len(path)-1]
		// 不选这个数，继续往下递归，考察下一个数
		dfs(start+1,path)
	}
	dfs(0,[]int{})
	return res
}

/*
在执行子递归之前，加入解集，即，在递归压栈前 “做事情”
 */
func subsets2(nums []int) [][]int {
	res := [][]int{}
	var dfs func(startIndex int,path []int)
	dfs = func(startIndex int, path []int) {
		temp := make([]int,len(path))
		copy(temp,path)
		res = append(res,temp)
		for i := startIndex; i < len(nums); i++ {
			path = append(path,nums[i])
			// 基于选这个数，继续递归，传入的是i+1，不是startIndex+1
			dfs(i+1,path)
			path = path[:len(path)-1]
		}
	}
	dfs(0,[]int{})
	return res
}