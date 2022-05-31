package main

import "fmt"

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 */
func main() {
	nums := []int{1,2,3}
	res := permute(nums)
	fmt.Println(res)
}
func permute(nums []int) [][]int {
	res := [][]int{}
	//记录使用过的数
	used := map[int]bool{}
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int,len(path))
			copy(temp,path)
			res = append(res,temp)
		}
		for _, n := range nums {
			if used[n] {
				continue
			}
			path = append(path,n)
			//记录一下当前使用过的数
			used[n] = true
			dfs(path)
			path = path[:len(path)-1]
			//撤销记录
			used[n]=false
		}
	}
	dfs([]int{})
	return res
}