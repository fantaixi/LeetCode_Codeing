package _49

/*
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
 */
func intersection(nums1 []int, nums2 []int) []int {
	ans := []int{}
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] = 1
	}
	for _, v := range nums2 {
		// 利用count>0，实现重复值只拿一次放入返回结果中
		if count, ok := m[v]; ok && count > 0 {
			ans = append(ans,v)
			m[v]--
		}
	}
	return ans
}