package main

func main() {

}
/**
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val和next。val是当前节点的值，next是指向下一个节点的指针/引用。
如果要使用双向链表，则还需要一个属性prev以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第index个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为val的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第index个节点之前添加值为val 的节点。如果index等于链表的长度，则该节点将附加到链表的末尾。
如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引index 有效，则删除链表中的第index 个节点。

 */

/*
// 节点结构体
type Node struct {
	Val int
	Next *Node
}
// MylinkedList是一个对象，需要去实现LinkedList接口的所有方法
type MyLinkedList struct {
	DummyHead *Node //虚拟头节点
	Size int //链表长度
}

//创建一个链表
func Constructor() MyLinkedList { //成功
	// 创建一个虚拟头节点，非真正的头节点(真正的头节点是数组第一个节点)
	dummyHead := &Node{
		Val : -1,
		Next : nil,
	}
	return MyLinkedList{DummyHead: dummyHead,Size: 0}
}
// 获取index位置的元素Val
// 获取到第index个节点数值，如果index是非法数值直接返回-1， 注意index是从0开始的，第0个节点就是头结点
func (this *MyLinkedList) Get(index int) int {
	// 判断index是否非法
	if (index < 0) || (index > (this.Size - 1)) {
		return -1
	}
	// 查找
	var cur *Node = this.DummyHead// dummy节点的index = -1
	for i := 0;i <= index;i++ {//找到index为 index的节点
		cur = cur.Next //0,1,2,3,4....index
	}
	return cur.Val
}

// 在头节点前面再次添加一个节点
func (this *MyLinkedList) AddAtHead(val int)  { //成功
	// 在dummy节点后面直接添加一个节点
	var newNode *Node = &Node{Val:val, Next:nil,} //所有变量都要显示的初始化
	newNode.Next = this.DummyHead.Next
	this.DummyHead.Next = newNode
	this.Size++
}


// 在尾结点添加节点
func (this *MyLinkedList) AddAtTail(val int)  {
	var newNode *Node = &Node{val, nil}
	cur := this.DummyHead
	for cur.Next != nil { //找到末尾节点
		cur = cur.Next
	}
	cur.Next = newNode //新元素添加到末尾节点后面
	this.Size++
}

// 在index节点前面添加一个节点
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.Size {
		return
	}else if index == this.Size { //直接添加到末尾
		this.AddAtTail(val)
		return
	}else if index < 0 {
		index = 0
	}
	var newNode *Node = &Node{val, nil}
	cur := this.DummyHead
	for i:=0; i<=index-1; i++ { //找到index为 index-1的节点,如果index=0，则不会进入循环，直接插入到dummy后面
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	this.Size++
}
// 删除index节点
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	// 判断是否有效
	if index > this.Size-1 || index < 0 {
		return
	}
	cur := this.DummyHead
	for i := 0; i <= index-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	this.Size--
}

*/