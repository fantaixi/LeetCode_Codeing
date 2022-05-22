package _5

import "sort"

/*
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/
/*
思路
外层循环：指针 i 遍历数组。
内层循环：用双指针，去寻找满足三数之和 == 0 的元素

先排序的意义
便于跳过重复元素，如果当前元素和前一个元素相同，跳过。

双指针的移动时，避免出现重复解
找到一个解后，左右指针同时向内收缩，为了避免指向重复的元素，需要：

左指针在保证left < right的前提下，一直右移，直到指向不重复的元素
右指针在保证left < right的前提下，一直左移，直到指向不重复的元素
小优化
排序后，如果外层遍历的数已经大于0，则另外两个数一定大于0，sum不会等于0，直接break。
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		//去重
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			n2, n3 := nums[left], nums[right]
			if n1+n2+n3 == 0 {
				res = append(res,[]int{n1,n2,n3})
				for left < right && nums[left] == n2 {  //去重
					left++
				}
				for left<right&&nums[right]==n3 {  //去重
					right--
				}
			}else if n1+n2+n3 < 0{
				left++
			}else {
				right--
			}
		}
	}
	return res
}
