package _04

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定二叉树的根节点 root ，返回所有左叶子之和。
 */
/*
判断当前节点是不是左叶子是无法判断的，必须要通过节点的父节点来判断其左孩子是不是左叶子。
如果该节点的左节点不为空，该节点的左节点的左节点为空，该节点的左节点的右节点为空，则找到了一个左叶子
 */
func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	findLeft(root,&res)
	return res
}
/*
当遇到左叶子节点的时候，记录数值，然后通过递归求取左子树左叶子之和，和右子树左叶子之和，相加便是整个树的左叶子之和
 */
func findLeft(root *TreeNode, res *int) {
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*res += root.Left.Val
	}
	if root.Left != nil {
		findLeft(root.Left,res)
	}
	if root.Right != nil {
		findLeft(root.Right,res)
	}

}

/*
迭代
 */
func sumOfLeftLeaves1(root *TreeNode) int {
	res := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
				res += node.Left.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}