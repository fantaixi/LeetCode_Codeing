package _42

type ListNode struct {
	Val  int
	Next *ListNode
}
/*
给定一个链表的头节点 head，返回链表开始入环的第一个节点。如果链表无环，则返回null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。
注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。
 */
/*
1、判断链表是否环
让fast走两步，slow走一步，如果有环的话，那它们一定会在环中相遇：fast一定会先进环，而且每次都是走两步，
所以在环中的情况是一步一步追赶slow，最终一定会相遇
2、如果有环，如何找到这个环的入口
假设头结点到入口结点数为x，fast和slow相遇结点数为y，从相遇结点到入口结点数为z，则有：
相遇时：slow：x+y ， fast ：x+y+n（y+z）   （n（y+z）：表示走了n圈 ）
因为fast是slow的两倍，所以：2（x+y） = x+y+n（y+z）
同时消去一个x+y： x+y = n（y+z）
y移到右边，同时提取一个（y+z） ： x=（y+z）（n-1）+z
因为n一定大于等于1，所以当n=1的时候，x = z，
所以就有：从头结点出发一个指针，从相遇节点 也出发一个指针，这两个指针每次只走一个节点， 那么当这两个指针相遇的时候就是 环形入口的节点。（这次只走一步）
当n大于1的时候，结果也和n=1的情况一样
 */func detectCycle(head *ListNode) *ListNode {
	fast,slow := head,head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}

