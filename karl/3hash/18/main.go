package _8

import "sort"

/*
给你一个由 n 个整数组成的数组nums ，和一个目标值 target 。
请你找出并返回满足下述全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]]（若两个四元组元素一一对应，则认为两个四元组重复）：
0 <= a, b, c, d< n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。
 */
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i+1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			left,right := j+1,len(nums)-1
			for left < right {
				n3,n4 := nums[left],nums[right]
				sum := n1+n2+n3+n4
				if sum < target {
					left++
				}else if sum > target {
					right--
				}else {
					res = append(res,[]int{n1,n2,n3,n4})
					for left < right && n3 == nums[left+1] {
						left++
					}
					for left < right && n4 == nums[right-1] {
						right--
					}
					right--
					left++
				}
			}
		}
	}
	return res
}
