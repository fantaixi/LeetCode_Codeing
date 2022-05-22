package _8

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。
假设一个二叉搜索树具有如下特征：
节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
 */
/*
中序遍历：遍历出来的数组是有序的
 */
func isValidBST(root *TreeNode) bool {
	// 保存上一个指针
	var prev *TreeNode
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		leftRes := dfs(node.Left)
		// 当前值小于等于前一个节点的值，返回false
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		rightRes := dfs(node.Right)
		return leftRes && rightRes
	}
	return dfs(root)
}

/*
递归
 */
func isValidBST1(root *TreeNode) bool {
	return dfs(root,math.MinInt64,math.MaxInt64)
}
func dfs(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return dfs(node.Left,min,node.Val) && dfs(node.Right,node.Val,max)
}