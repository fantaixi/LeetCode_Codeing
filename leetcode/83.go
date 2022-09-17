package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil  {
		return nil
	}
	newHead := head
	for newHead.Next != nil {
		if newHead.Val == newHead.Next.Val {
			newHead.Next = newHead.Next.Next
		}else {
			newHead = newHead.Next
		}
	}
	return head
}

/*
使用递归
 */
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates1(head.Next)
	if head.Val == head.Next.Val {
		head = head.Next
	}
	return head
}
