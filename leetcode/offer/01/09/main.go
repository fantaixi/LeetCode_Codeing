package main

import "container/list"

func main() {

}
type CQueue struct {
	stack1,stack2 *list.List
}


func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.stack1.PushBack(value)
}


func (this *CQueue) DeleteHead() int {
	//如果stack2为空就把stack1的元素全部倒置放入stack2中并在stack1中删除
	if this.stack2.Len() == 0 {
		for this.stack1.Len() != 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	//不为空就返回stack2的Back并删除
	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */