package main

import "sort"

/*
给你一个整数数组 arr，请你检查是否存在两个整数 N 和 M，满足 N 是 M 的两倍（即，N = 2 * M）。
*/
func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		left, right, res := 0, len(arr)-1, arr[i]*2
		//处理边界
		if arr[i] >= 0 {
			left = i + 1
		} else {
			right = i - 1
		}
		//双指针
		for left <= right {
			mid := left + (right-left)>>1
			if arr[mid] == res {
				return true
			} else if arr[mid] > res {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}
