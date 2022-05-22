package main

import "container/list"

func main() {

}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
 */
/*
使用前序求的就是深度，使用后序求的是高度。
先求它的左子树的深度，再求右子树的深度，
最后取左右深度最大的数值 再+1 （加1是因为算上当前中间节点）就是目前节点为根节点的树的深度。
*/
/*
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}else {
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		if left > right {
			return left+1
		}
		return right+1
	}
}
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left),maxDepth(root.Right))+1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
层次遍历
 */
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0{
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans++
	}
	return ans
}
