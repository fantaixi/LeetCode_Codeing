package _35
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 */
/*
二叉搜索树性质决定了：如果 p.val 和 q.val 都比 root.val 小，则p、q肯定在 root 的左子树。
那问题规模就变小了，递归左子树就行！
如果 p.val 和 q.val 都比 root.val 大，递归右子树就行！
其他情况，root 即为所求！那么简单吗？为什么？
只要 p.val 和 q.val 不是都大于(小于) root.val，即只要 p, q 不同处在 root 的一个子树
就只有这三种情况：
p 和 q 分居 root 的左、右子树。
root 就是 p，q 在 p 的子树中。
root 就是 q，p 在 q 的子树中
而这三种情况，p 和 q 的最近公共祖先都是 root
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left,p,q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right,p,q)
	}
	return root
}

/*
迭代
开启 while 循环，当 root 为 null 时就结束循环（root 就是一个指针）。
如果 p.val、q.val 都小于 root.val，它们都在 root 的左子树，root=root.left，遍历到 root 的左子节点。
如果 p.val、q.val 都大于 root.val，它们都在 root 的右子树，root=root.right，遍历到 root 的右子节点。
其他情况，当前的 root 就是最近公共祖先，结束遍历， break。
返回 root，即，break 时的 root 节点就是最近公共祖先
 */
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		}else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		}else {
			break
		}
	}
	return root
}