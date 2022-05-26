package _0

import "sort"

/*
给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
candidates中的每个数字在每个组合中只能使用一次。
注意：解集不能包含重复的组合
 */
func combinationSum2(candidates []int, target int) [][]int {
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
			if i-1 >= startIndex && candidates[i-1] == candidates[i] {
				continue
			}
			temp = append(temp,candidates[i])
			backstrack(i+1,temp,sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	backstrack(0,[]int{},0)
	return res
}
