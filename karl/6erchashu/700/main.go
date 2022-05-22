package _00
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定二叉搜索树（BST）的根节点root和一个整数值val。
你需要在 BST 中找到节点值等于val的节点。 返回以该节点为根的子树。 如果节点不存在，则返回null。
 */
/*
二叉搜索树是一个有序树：
若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
它的左、右子树也分别为二叉搜索树
 */
/*
如果root为空，或者找到这个数值了，就返回root节点。
因为二叉搜索树的节点是有序的，所以可以有方向的去搜索。
如果root->val > val，搜索左子树，如果root->val < val，就搜索右子树，最后如果都没有搜索到，就返回NULL。
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left,val)
	}
	if root.Val < val {
		return searchBST(root.Right,val)
	}
	return nil
}

/*
迭代
对于二叉搜索树，不需要回溯的过程，因为节点的有序性就帮我们确定了搜索的方向
 */
func searchBST1(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val {
			root = root.Left
		}else if root.Val < val {
			root = root.Right
		}else {
			return root
		}
	}
	return nil
}