package main

func main() {

}
 type ListNode struct {
     Val int
     Next *ListNode
 }

func reversePrint1(head *ListNode) []int {
    if head == nil {
        return nil
    }
    return apeendData(head)
}
func apeendData(head *ListNode) []int {
    if head.Next != nil {
        list := apeendData(head.Next)
        list = append(list,head.Val)
        return list
    }
    return []int{head.Val}
}

func reversePrint(head *ListNode) []int {
//    原地反转链表
//    var newNode *ListNode
//    for head != nil {
//        next := head.Next
//        head.Next = newNode
//        newNode = head
//        head = next
//    }
//    res := []int{}
//    for newNode != nil {
//        res = append(res,newNode.Val)
//        newNode = newNode.Next
//    }
//    return res

    //遍历链表存入数组，再反转数组
    //res := []int{}
    //if head == nil {
    //    return nil
    //}
    //for head != nil {
    //    res = append(res,head.Val)
    //    head = head.Next
    //}
    //reverse(res)
    //return res

    //遍历链表，记录链表长度,数组按照下标从大到小存储结点
    //if head == nil {
    //    return nil
    //}
    //num := 0
    //newHead := head
    //for newHead != nil {
    //    num++
    //    newHead = newHead.Next
    //}
    //res := make([]int,num)
    //for ; num > 0; num-- {
    //    res[num-1] = head.Val
    //    head = head.Next
    //}
    //return res

    //用栈
    temp := []int{}
    res := []int{}
    if head == nil {
        return nil
    }
    for head != nil {
        temp = append(temp,head.Val)
        head = head.Next
    }
    for len(temp) > 0 {
        res = append(res,temp[len(temp)-1])
        temp = temp[:len(temp)-1]
    }
    return res
}
func reverse(list []int) []int {
    n := len(list)
    for i := 0; i < n/2; i++ {
        list[i],list[n-i-1] = list[n-i-1],list[i]
    }
    return list
}