package _54
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
给定一个不重复的整数数组nums 。最大二叉树可以用下面的算法从nums 递归地构建:
创建一个根节点，其值为nums 中的最大值。
递归地在最大值左边的子数组前缀上构建左子树。
递归地在最大值 右边 的子数组后缀上构建右子树。
返回nums 构建的 最大二叉树 。
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	index := findMax(nums)
	root := &TreeNode{
		Val: nums[index],
		Left: constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
	return root
}

func findMax(nums []int) (index int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return index
}