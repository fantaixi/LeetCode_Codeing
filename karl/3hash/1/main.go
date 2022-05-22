package _1

/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
 */

func twoSum(nums []int, target int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res,i,j)
			}
		}
	}
	return res
}

func twoSum1(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{index,i}
		}else {
			m[v] = i
		}
	}
	return nil
}
