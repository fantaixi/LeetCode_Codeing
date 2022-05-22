package _32
/*
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
实现 MyQueue 类：
void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
说明：
你 只能 使用标准的栈操作 —— 也就是只有push to top,peek/pop from top,size, 和is empty操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
 */
type MyQueue struct {
	// 双栈实现队列
	instack []int  // push
	outStack []int  // pop || peek
}


func Constructor() MyQueue {
	return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
	// 压入尾部即可
	this.instack = append(this.instack,x)
}

// 当进行pop或peek操作时，如果outstack空，都需要将instack中的元素全部转入outstack中
func (this *MyQueue) in2out() {
	// 当outstack空时，需要将instack全部转入outstack
	for len(this.instack) > 0 {
		this.outStack = append(this.outStack,this.instack[len(this.instack)-1])
		this.instack = this.instack[:len(this.instack)-1]
	}
}

func (this *MyQueue) Pop() int {
	// pop时需要先判断outstack是否为空
	if len(this.outStack) == 0 {
		this.in2out()
	}
	// pop时出栈top, 然后更新outstack
	val := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return val
}


func (this *MyQueue) Peek() int {
	// 因为outstack才是出栈功能部分，需要先判断是否为空
	if len(this.outStack) == 0 {
		this.in2out()
	}
	return this.outStack[len(this.outStack)-1]
}


func (this *MyQueue) Empty() bool {
	// 仅当入栈和出栈都为空时，queue才为空
	return len(this.outStack)==0 && len(this.instack)==0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
