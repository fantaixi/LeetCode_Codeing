package _00

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return dfs(p,q)
}

func dfs(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return dfs(left.Left,right.Left) && dfs(left.Right,right.Right)
}
