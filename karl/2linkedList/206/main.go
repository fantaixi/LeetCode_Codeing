package _06

type ListNode struct {
	Val  int
	Next *ListNode
}
/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/
//双指针
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
//递归
func reverseList1(head *ListNode) *ListNode {
	return reverse(nil,head)
}
func reverse(pre, head *ListNode) *ListNode {
	if head ==	nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return reverse(head,next)
}
