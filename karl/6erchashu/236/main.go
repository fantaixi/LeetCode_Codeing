package _36
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，
最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 */
/*
对于根节点 root，p、q 的分布，有两种可能：
p、q 分居 root 的左右子树，则 LCA 为 root。
p、q 存在于 root 的同一侧子树中，就变成规模小一点的相同问题。

递归函数：返回当前子树中 p 和 q 的 LCA。如果没有 LCA，就返回 null。
从根节点 root 开始往下递归遍历……
如果遍历到 p 或 q，比方说 p，则 LCA 要么是当前的 p（q 在 p 的子树中），
要么是 p 之上的节点（q 不在 p 的子树中），不可能是 p 之下的节点，不用继续往下走，返回当前的 p。
当遍历到 null 节点，空树不存在 p 和 q，没有 LCA，返回 null。
当遍历的节点 root 不是 p 或 q 或 null，则递归搜寻 root 的左右子树：
如果左右子树的递归都有结果，说明 p 和 q 分居 root 的左右子树，返回 root。
如果只是一个子树递归调用有结果，说明 p 和 q 都在这个子树，返回该子树递归结果。
如果两个子树递归结果都为 null，说明 p 和 q 都不在这俩子树中，返回 null。
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left,p,q)
	right := lowestCommonAncestor(root.Right,p,q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}