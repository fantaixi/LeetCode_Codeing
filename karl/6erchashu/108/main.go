package _08

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
*/
/*
取数组中间值为根节点
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	left, right := 0, len(nums)
	midLength := left + (right-left)>>1
	root := &TreeNode{nums[midLength], nil, nil}
	root.Left = sortedArrayToBST(nums[:midLength])
	root.Right = sortedArrayToBST(nums[midLength+1:])
	return root
}
