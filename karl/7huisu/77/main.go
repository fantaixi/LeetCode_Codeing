package _77

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。
*/
var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	backstack(n, k, 1, []int{})
	return res
}

func backstack(n, k, startIndex int, path []int) {
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		res = append(res, temp)
		//return
	}
	//修剪优化：
	//如果for循环选择的起始位置之后的元素个数已经不足需要的元素个数了，那么就没有必要搜索
	/*/
	优化过程如下：
	已经选择的元素个数：len(path);
	还需要的元素个数为: k - len(path);
	在集合n中至多要从该起始位置 : n - (k - len(path)) + 1，开始遍历
	为什么有个+1呢，因为包括起始位置，需要的是一个左闭的集合。
	*/
	if len(path)+n-startIndex+1 < k {
		return
	}
	for i := startIndex; i <= n; i++ {
		path = append(path, i)     // 处理节点
		backstack(n, k, i+1, path) //递归：控制树的纵向遍历，注意下一层搜索要从i+1开始
		path = path[:len(path)-1]  // 回溯，撤销处理的节点
	}
}

//换一种写法
func combine1(n int, k int) [][]int {
	res := [][]int{}
	var helper func(start int, path []int)
	helper = func(start int, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := start; i <= n-(k-len(path))+1; i++ { //i <= n-(k-len(path))  优化的部分
			path = append(path, i)
			helper(i+1, path)
			path = path[:len(path)-1]
		}
	}
	helper(1, []int{})
	return res
}
