package main

func main() {

}
type TreeNode struct {
	    Val int
	    Left *TreeNode
	     Right *TreeNode
	 }
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countNodes(root.Left)
	right := countNodes(root.Right)
	return left+right+1
}