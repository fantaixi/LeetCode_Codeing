package _38

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），
使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
提醒一下，二叉搜索树满足下列约束条件：
节点的左子树仅包含键 小于 节点键的节点。 节点的右子树仅包含键 大于 节点键的节点。 左右子树也必须是二叉搜索树。
*/
/*
一个有序数组[2, 5, 13]，求从后到前的累加数组，也就是[20, 18, 13]
累加的顺序是右中左，反中序遍历这个二叉树，然后顺序累加
*/
func convertBST(root *TreeNode) *TreeNode {
	pre := 0
	dfs(root, &pre)
	return root
}
func dfs(node *TreeNode, pre *int) *TreeNode {
	if node == nil {
		return nil
	}
	//先遍历右边
	dfs(node.Right, pre)
	//暂存总和的值
	temp := *pre
	//变更总和值
	*pre += node.Val
	//更新节点值
	node.Val += temp
	//遍历左节点
	dfs(node.Left, pre)
	return node
}
