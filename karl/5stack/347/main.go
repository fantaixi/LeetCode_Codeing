package _47

import "sort"

/*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
 */
func topKFrequent(nums []int, k int) []int {
	res := []int{}
	num_map := map[int]int{}
	for _, val := range nums {
		num_map[val]++
	}
	for key, _ := range num_map {
		res = append(res,key)
	}
	sort.Slice(res, func(i, j int) bool {
		return num_map[res[i]] > num_map[res[j]]
	})
	return res[:k]
}
