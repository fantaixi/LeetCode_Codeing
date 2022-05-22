package _25new

/*
请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
实现 MyStack 类：
void push(int x) 将元素 x 压入栈顶。
int pop() 移除并返回栈顶元素。
int top() 返回栈顶元素。
boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。
注意：
你只能使用队列的基本操作 —— 也就是push to back、peek/pop from front、size 和is empty这些操作。
你所使用的语言也许不支持队列。你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列, 只要是标准的队列操作即可。
 */
/*
使用一个队列模拟
 */
type MyStack struct {
	queue []int
}


func Constructor() MyStack {
	return MyStack{}
}


func (this *MyStack) Push(x int)  {
	this.queue = append(this.queue,x)
}


func (this *MyStack) Pop() int {
	n := len(this.queue)-1
	for n != 0 {  //除了最后一个，其余的都重新添加到队列里
		val := this.queue[0]
		this.queue=this.queue[1:]
		this.queue=append(this.queue,val)
		n--
	}
	//弹出元素
	val := this.queue[0]
	this.queue=this.queue[1:]
	return val
}


func (this *MyStack) Top() int {
	//利用Pop函数，弹出来的元素重新添加
	val := this.Pop()
	this.queue = append(this.queue,val)
	return val
}


func (this *MyStack) Empty() bool {
	return len(this.queue)==0
}