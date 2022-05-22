package _30

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
差值是一个正数，其数值等于两值之差的绝对值。
 */
/*
中序遍历
 */
func getMinimumDifference(root *TreeNode) int {
	var res []int
	findMin(root,&res)
	min := math.MaxInt64
	for i := 1; i < len(res); i++ {
		temp := res[i] - res[i-1]
		if temp < min {
			min = temp
		}
	}
	return min
}
func findMin(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	findMin(root.Left,res)
	*res = append(*res,root.Val)
	findMin(root.Right,res)
}

/*
中序遍历的时候求最小值
 */
func getMinimumDifference1(root *TreeNode) int {
	var prev *TreeNode
	min := math.MaxInt64
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil && node.Val-prev.Val < min {
			min = node.Val - prev.Val
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)
	return min
}