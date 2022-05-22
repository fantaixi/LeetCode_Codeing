package _50

import "sort"

/*
给你两个整数数组nums1 和 nums2 ，请你以数组形式返回两数组的交集。
返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。
可以不考虑输出结果的顺序。
*/
//hash
func intersect(nums1 []int, nums2 []int) []int {
	ans := []int{}
	numCnt := map[int]int{}
	for i := range nums1 {
		numCnt[nums1[i]]++
	}
	for i := range nums2 {
		if count, ok := numCnt[nums2[i]]; ok && count > 0 {
			ans = append(ans, nums2[i])
			numCnt[nums2[i]]--
		}
	}
	return ans
}
//换个写法
func intersect1(nums1 []int, nums2 []int) []int {
	//先存长度小的的数组，避免浪费空间
	if len(nums1) > len(nums2) {
		return intersect(nums2,nums1)
	}
	numCnt := map[int]int{}
	for _, v := range nums1 {
		numCnt[v]++
	}
	ans := []int{}
	for _, v := range nums2 {
		if numCnt[v] > 0 {
			ans = append(ans,v)
			numCnt[v]--
		}
	}
	return ans
}

//排序+双指针
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1,length2,index1,index2 := len(nums1),len(nums2),0,0
	ans := []int{}
	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1++
		}else if nums1[index1] > nums2[index2] {
			index2++
		}else {
			ans = append(ans,nums1[index1])
			index1++
			index2++
		}
	}
	return ans
}