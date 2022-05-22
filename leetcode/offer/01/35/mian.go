package main

func main() {

}

type Node struct {
	Val int
	Next *Node
	Random *Node
}
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 1. 新生成的节点放到原节点的后面，例如1->2->3这样的链表就变成了这样1->1'->2->2'->3->3'
	//此时拷贝节点的random指针都指向nil
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val,Next: node.Next}
	}
	//cur.Next指向的替身替身1，cur.Random指向的原节点的random，
	//那么cur.Random.Next指向的是谁？答案是cur.Random的替身2
	//那么将 替身1的random 指向 替身2，就完成了copy节点的random指向，非常巧妙
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil { // random指针可能为空
			node.Next.Random = node.Random.Next  //cur.Next是替身1，cur.Random.Next是替身2
		}
	}
	//拆分原节点 和 copy节点
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	// map中存的是（原节点->新节点）的映射关系，此时新节点只有val，指针并没有安排上
	m := make(map[*Node]*Node)
	for cur := head; cur != nil; cur = cur.Next {
		m[cur] = &Node{
			Val: cur.Val,
		}
	}
	//将新节点串起来，组成新链表
	for cur := head; cur != nil; cur = cur.Next {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
	}
	return m[head]
}