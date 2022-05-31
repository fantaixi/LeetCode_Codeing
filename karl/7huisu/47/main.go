package _7

import "sort"

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 */
/*
产生重复排列的原因：
比如[1,1,2]，先选第一个1，和先选第二个1，往后的情况是一样的。
即，“同一层” 的选项出现重复，或者说，当前可选的选项出现重复。
 */
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	//为了方便在迭代中识别出重复，先对 nums 中数字升序排序，使得相同的数字相邻
	sort.Ints(nums)
	used := make([]bool,len(nums))
	dfs([]int{},nums,used,&res)
	return res
}
func dfs(path, nums []int, used []bool, res *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int,len(nums))
		copy(temp,path)
		*res = append(*res,temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		//使用过了就不再使用
		if used[i] {
			continue
		}
		/*
		如果当前的选项nums[i]，与同一层的前一个选项nums[i-1]相同，且nums[i-1]存在，且没有被使用过，则忽略选项nums[i]。
		如果nums[i-1]被使用过，它会被第一条修剪掉，不是选项了，即便它和nums[i]重复，nums[i]还是可以选的。
		 */
		if i-1 >= 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}
		path = append(path,nums[i])
		used[i] = true
		dfs(path,nums,used,res)
		path = path[0:len(path)-1]
		used[i] = false
	}
}