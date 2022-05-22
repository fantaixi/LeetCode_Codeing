package mianshi0207
type ListNode struct {
	Val  int
	Next *ListNode
}
//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

/*
让两个指针同时遍历两个链表，假设链表headA的结点数量为a，链表headB的结点数量为b，公共结点数量为c，那么：
链表A头结点到公共结点之前有a-c个结点
链表B头结点到公共结点之前有b-c个结点
指针A遍历完链表A之后，再开始遍历链表B，走到公共结点的时候，一共走了  a+(b-c) 步
指针B遍历完链表B之后，再开始遍历链表A，走到公共结点的时候，一共走了  b+(a-c) 步
如果有公共结点：那么就有  a+(b-c) = b+(a-c)
如果c > 0 ,那么此时c对应的结点就是公共结点，
如果c = 0 ,那么此时a b 同时指向 nil
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointA,pointB := headA,headB
	for pointA != pointB {
		if pointA != nil {
			pointA = pointA.Next
		}else {
			pointA = headB
		}
		if pointB != nil {
			pointB = pointB.Next
		}else {
			pointB = headA
		}
	}
	return pointA
}
