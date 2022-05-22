package _25
/*
使用两个队列模拟
 */
type MyStack struct {
	queue1 []int
	queue2 []int
}


func Constructor() MyStack {
	return MyStack{}
}


func (this *MyStack) Push(x int)  {
	this.queue2 = append(this.queue2,x)
	this.Move()
}

func (this *MyStack) Move() {
	if len(this.queue1) == 0 {
		this.queue1,this.queue2 = this.queue2,this.queue1
	}else {
		this.queue2 = append(this.queue2,this.queue1[0])
		this.queue1 = this.queue1[1:]	//去除第一个元素
		this.Move()     //重复
	}
}

func (this *MyStack) Pop() int {
	val := this.queue1[0]
	this.queue1 = this.queue1[1:]	//去除第一个元素
	return val
}


func (this *MyStack) Top() int {
	return this.queue1[0]	//直接返回
}


func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}
