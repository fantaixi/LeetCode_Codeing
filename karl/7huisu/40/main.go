package _0

import "sort"

/*
给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
candidates中的每个数字在每个组合中只能使用一次。
注意：解集不能包含重复的组合
 */
func combinationSum2(candidates []int, target int) [][]int {
	//给定的数组可能有重复的元素，先排序，使得重复的数字相邻，方便去重
	sort.Ints(candidates)
	res := [][]int{}
	var backstrack func(startIndex int,temp []int,sum int)
	backstrack = func(startIndex int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				t := make([]int,len(temp))
				copy(t,temp)
				res = append(res,t)
			}
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			//忽略掉同一层重复的选项，避免产生重复的组合
			if i-1 >= startIndex && candidates[i-1] == candidates[i] {
				continue
			}
			temp = append(temp,candidates[i])
			//当前选择的数字不能和下一个选择的数字重复，给子递归传i+1，避免与当前选的i重复。
			backstrack(i+1,temp,sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	backstrack(0,[]int{},0)
	return res
}


func combinationSum(candidates []int, target int) [][]int {
	var path []int
	var res [][]int
	var history map[int]bool
	history = make(map[int]bool)
	sort.Ints(candidates)
	backtrack(0,0,target,candidates,path,&res,history)
	return res
}
func backtrack(startIndex, sum, target int, candidates, path []int, res *[][]int, history map[int]bool) {
	if sum == target {
		temp := make([]int,len(path))
		copy(temp,path)
		*res = append(*res,temp)
	}
	if sum > target {
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		// history[i - 1] == true，说明同一树枝candidates[i - 1]使用过
		// history[i - 1] == false，说明同一树层candidates[i - 1]使用过
		if i > 0 && candidates[i] == candidates[i-1] && history[i-1] == false {
			continue
		}
		path = append(path,candidates[i])
		sum += candidates[i]
		history[i] = true
		backtrack(i+1,sum,target,candidates,path,res,history)
		path = path[:len(path)-1]
		sum -= candidates[i]
		history[i] = false
	}
}