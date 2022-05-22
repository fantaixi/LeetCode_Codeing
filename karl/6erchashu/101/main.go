package main

func main() {

}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定一个二叉树，检查它是否是镜像对称的。
 */
/*
//使用递归
func isSymmetric(root *TreeNode) bool {
	return dfs(root.Left,root.Right)
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
	return dfs(left.Left,right.Right) && dfs(left.Right,right.Left)
}
*/

//使用队列
func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue,root.Left,root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right ==nil {
			continue
		}
		if left == nil || right==nil || left.Val != right.Val {
			return false
		}
		queue = append(queue,left.Left,right.Right,right.Left,left.Right)
	}
	return true
}