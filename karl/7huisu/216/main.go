package _16
/*
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
说明：
所有数字都是正整数。
解集不能包含重复的组合。
示例 1: 输入: k = 3, n = 7 输出: [[1,2,4]]
示例 2: 输入: k = 3, n = 9 输出: [[1,2,6], [1,3,5], [2,3,4]]
 */
func combinationSum3(k int, n int) [][]int {
	var path []int
	var res [][]int
	backstack(n,k,1,&path,&res)
	return res
}

func backstack(n, k, startIndex int, path *[]int,res *[][]int) {
	if len(*path) == k {
		sum := 0
		temp := make([]int,k)
		//将结果放入path，避免使用重复数字
		for k, v := range *path {
			sum += v
			temp[k] = v
		}
		if sum == n {
			*res = append(*res,temp)
		}
		return
	}
	for i := startIndex; i <= 9-(k-len(*path))+1 ; i++ {
		*path = append(*path,i)
		backstack(n,k,i+1,path,res)
		*path = (*path)[:len(*path)-1]
	}
}