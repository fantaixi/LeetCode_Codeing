package _9
type ListNode struct {
	Val  int
	Next *ListNode
}
/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 */
//使用双指针
/*
需要找到倒数第n节点前面的节点
设置slow, fast 指针， slow永远保持距离fast n 步（如保持2步：slow.Next.Next == fast）。
当fast到达末尾节点， slow就是倒数第n+1个。 删除第n个节点

 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	fast,slow := dummyHead,dummyHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow,fast = slow.Next,fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
//换个写法
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	fast := head
	slow := dummyHead
	i := 1
	for fast != nil {
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}
		i++
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}