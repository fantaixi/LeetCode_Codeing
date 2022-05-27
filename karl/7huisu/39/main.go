package _9
/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。
说明：
所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
 */
func combinationSum(candidates []int, target int) [][]int {
	var path []int
	var res [][]int
	backtrack(0,0,target,candidates,path,&res)
	return res
}
func backtrack(startIndex, sum, target int, candidates, path []int, res *[][]int) {
	if sum == target {
		temp := make([]int,len(path))
		copy(temp,path)
		*res = append(*res,temp)
		return
	}
	if sum > target {
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		path = append(path,candidates[i])
		sum += candidates[i]
		backtrack(i,sum,target,candidates,path,res)
		path = path[:len(path)-1]
		sum -= candidates[i]
	}
}

/*
回溯：在包含问题所有解的空间树中，用DFS的方式，从根节点出发，搜索整棵解空间树。
搜索至任何一个节点时，总是会先判断当前节点是否可以通往最后的合法解。
如果不可以，则结束对「以当前节点为根节点的子树」的搜索，向父节点回溯，回到之前的状态，搜索下一个分支。
否则，进入该子树，继续以DFS的方式搜索。
空间树中的节点是动态的，即，当前有哪些选项可选择，是根据上一步的选择得出的，所以做回溯时，要把状态还原成进入当前节点之前的状态。
确定出问题的解空间树，它是隐式的，不是显式的一棵树。不熟练的就画图看看。
然后，明确每个节点的扩展搜索规则。
然后进行DFS搜索，并注意剪枝，避免无效的搜索。
 */
func combinationSum1(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func(start,sum int,temp []int)
	dfs = func(start, sum int, temp []int) {
		if sum >= target {
			if sum == target {
				newTemp := make([]int,len(temp))
				copy(newTemp,temp)
				res = append(res,newTemp)
			}
			return
		}
		/*
		不产生重复组合怎么限制（剪枝）？
		只要限制下一次选择的起点，是基于本次的选择，这样下一次就不会选到本次选择同层左边的数。即通过控制 for 遍历的起点，去掉会产生重复组合的选项
		 */
		// 枚举当前可选的数，从start开始
		for i := start; i < len(candidates); i++ {
			// 选这个数
			temp = append(temp,candidates[i])
			//子递归传了 i 而不是 i+1 ，因为元素可以重复选入集合，如果传 i+1 就不重复了
			dfs(i,sum+candidates[i],temp)
			// 撤销选择，回到选择candidates[i]之前的状态，继续尝试选同层右边的数
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0,0,[]int{})
	return res
}