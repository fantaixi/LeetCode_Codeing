package main

type Node struct {
	val  int
	next *Node
	prev *Node
}

type MyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func Constructor() MyLinkedList {
	p := &Node{}
	return MyLinkedList{p, p, 0}
}

// 带头节点的链表
// 获取链表中第 index 个节点的值。如果索引无效，则返回-1
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}
	p := this.head
	for i := 0; i <= index; i++ {
		p = p.next
	}
	return p.val
}

// 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	p := this.head
	q := &Node{val, p.next, p}
	if p.next != nil {
		p.next.prev = q
	} else {
		this.tail = q
	}
	p.next = q
	this.length++
}

// 将值为 val 的节点追加到链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int) {
	// 先找到链表的最后一个元素
	p := this.tail
	p.next = &Node{val, nil, p}
	this.tail = p.next
	this.length++
}

// 在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.length {
		this.AddAtTail(val)
		return
	}
	if index > this.length {
		return
	}

	p := this.head
	// 在第index个节点之前添加值，则需要找到第index个结点的位置
	for i := 0; i <= index; i++ {
		p = p.next
	}
	q := &Node{val, p, p.prev}
	p.prev.next = q
	p.prev = q
	this.length++
}

// 如果索引 index 有效，则删除链表中的第 index 个节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	}
	p := this.head
	// 删除第index个结点，则需要找到第index个结点
	for i := 0; i <= index; i++ {
		p = p.next
	}

	p.prev.next = p.next
	if p.next != nil {
		p.next.prev = p.prev
	} else {
		this.tail = p.prev
	}
	p.next = nil
	p.prev = nil
	this.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
