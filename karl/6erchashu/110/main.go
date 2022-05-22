package _10
/*
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
先看左右子树是否平衡，是则求得左右子树深度，差值大于1则不平衡
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 先看左右子树是否平衡，是则求得左右子树深度，差值大于1则不平衡
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	// 计算左右子树高度
	leftH,rightH := Depth(root.Left)+1,Depth(root.Right)+1
	if abs(leftH-rightH) > 1 {
		return false
	}
	return true
}

func Depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if Depth(root.Left) > Depth(root.Right) {
		return Depth(root.Left)+1
	}
	return Depth(root.Right)+1
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}