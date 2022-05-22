package _12
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给你二叉树的根节点root 和一个表示目标和的整数targetSum 。
判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。如果存在，返回 true ；否则，返回 false 。

叶子节点 是指没有子节点的节点。
 */
/*
不要去累加然后判断是否等于目标和，那么代码比较麻烦，可以用递减，让计数器count初始为目标和，然后每次减去遍历路径节点上的数值。
如果最后count == 0，同时到了叶子节点的话，说明找到了目标和。
如果遍历到了叶子节点，count不为0，就是没找到。
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}
	return hasPathSum(root.Left,targetSum) || hasPathSum(root.Right,targetSum)
}
/*
递归再精简一下
 */
func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum - root.Val == 0
	}
	return hasPathSum1(root.Left,targetSum - root.Val) || hasPathSum1(root.Right,targetSum-root.Val)
}