package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}
/*
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
*/
func removeElements(head *ListNode, val int) *ListNode {
	//设置一个虚拟头结点，方便用来删除头结点
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
