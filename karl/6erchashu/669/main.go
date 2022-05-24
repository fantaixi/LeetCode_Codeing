package _69

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。
通过修剪二叉搜索树，使得所有节点的值在[low, high]中。
修剪树 不应该改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在唯一的答案。
所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
*/
/*
递归：
如果root（当前节点）的元素小于low的数值，那么应该递归右子树，并返回右子树符合条件的头结点。
如果root(当前节点)的元素大于high的，那么应该递归左子树，并返回左子树符合条件的头结点。
接下来要将下一层处理完左子树的结果赋给root->left，处理完右子树的结果赋给root->right。
*/
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// 中序
	if root == nil {
		return nil
	}
	// 1.节点值小于low，递归root.Right，返回right
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	// 2.节点值大于high，递归root.Left，返回left
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	// 分别递归左右子树
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

/*
迭代
*/
func trimBST1(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 处理 root，让 root 移动到[low, high] 范围内，注意是左闭右闭
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	node := root
	// 此时 root 已经在[low, high] 范围内，处理左孩子元素小于 low 的情况（左节点是一定小于 root.Val，因此天然小于 high）
	for node != nil {
		for node.Left != nil && node.Left.Val < low {
			node.Left = node.Left.Right
		}
		node = node.Left
	}
	node = root
	// 此时 root 已经在[low, high] 范围内，处理右孩子大于 high 的情况
	for node != nil {
		for node.Right != nil && node.Right.Val > high {
			node.Right = node.Right.Left
		}
		node = node.Right
	}
	return root
}
