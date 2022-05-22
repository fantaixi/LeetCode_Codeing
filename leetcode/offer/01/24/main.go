package main

func main() {

}
type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for head.Next != nil {
		t := head.Next.Next
		//反转指针方向
		head.Next.Next = cur
		//将新头节点移到下一位
		cur = head.Next
		//连接断开的方向，重复上面的操作
		head.Next = t
	}
	return cur
}