package _50
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
一般来说，删除节点可分为两个步骤：
首先找到需要删除的节点； 如果找到了，删除它。 说明： 要求算法时间复杂度为 $O(h)$，h 为树的高度。
 */
/*
有以下五种情况：
第一种情况：没找到删除的节点，遍历到空节点直接返回了
找到删除的节点
第二种情况：左右孩子都为空（叶子节点），直接删除节点， 返回NULL为根节点
第三种情况：删除节点的左孩子为空，右孩子不为空，删除节点，右孩子补位，返回右孩子为根节点
第四种情况：删除节点的右孩子为空，左孩子不为空，删除节点，左孩子补位，返回左孩子为根节点
第五种情况：左右孩子节点都不为空，则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点。
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 1.空节点返回nil
	if root == nil {
		return nil
	}
	if root.Val == key {
		// 2.左右孩子都为空，直接删除节点返回nil
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 3.左空右不空，删除节点，右孩子补位
		if root.Left == nil && root.Right != nil {
			root = root.Right
			return root
		}
		// 4.左不空右空，删除节点，左孩子补位
		if root.Left != nil && root.Right == nil {
			root = root.Left
			return root
		}
		// 5.左右孩子节点都不为空，则将删除节点的左子树放到删除节点的右子树的最左面节点的左孩子的位置
		// root为删除节点，保存左孩子树
		left := root.Left
		// 保存右孩子树
		right := root.Right
		// 通过tmp不断遍历到root.Right的最左孩子节点
		temp := root.Right
		for temp.Left != nil {
			temp = temp.Left
		}
		// 右子树的最小节点，该节点大于root.Left整颗子树，left指向Left子树
		temp.Left = left
		// 更新root节点，相当于删除root节点
		root = right
		return root
	}
	if root.Val > key { // 递归左树
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key { // 递归右树
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
