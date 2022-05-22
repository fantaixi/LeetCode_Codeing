package _72

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。
如果存在，返回 true ；否则，返回 false 。
二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。
 */
/*
递归遍历 s 树的每个节点，看看：以当前节点为 root 的子树，是否和 t 树相同
从根节点开始，判断整个树是否和 t 相同，不相同则递归左右子树，是否和 t 相同
当遍历到 null 节点，始终没有返回true，则返回false
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if dfs(root, subRoot) {
		return true
	}
	return isSubtree(root.Left,subRoot) || isSubtree(root.Right,subRoot)
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
	return dfs(left.Left,right.Left)&&dfs(left.Right,right.Right)
}
