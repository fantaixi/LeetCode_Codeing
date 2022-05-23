package _01
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。
输入数据保证，新值和原始二叉搜索树中的任意节点值都不同。
注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回任意有效的结果。
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left,val)
	}else {
		root.Right = insertIntoBST(root.Right,val)
	}
	return root
}

/*
迭代
 */
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	node := root
	//记录前一个结点，用来赋值
	var prev *TreeNode
	for node != nil {
		if val < node.Val {
			prev = node
			node = node.Left
		}else {
			prev = node
			node = node.Right
		}
	}
	if val < prev.Val {
		prev.Left = &TreeNode{Val: val}
	}else {
		prev.Right = &TreeNode{Val: val}
	}
	return root
}
