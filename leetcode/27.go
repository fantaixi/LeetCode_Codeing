package main

/*
给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/
/*
右指针 right 指向当前将要处理的元素，左指针 left 指向下一个将要赋值的位置。
如果右指针指向的元素不等于 val，它一定是输出数组的一个元素，我们就将右指针指向的元素复制到左指针位置，然后将左右指针同时右移；
如果右指针指向的元素等于 val，它不能在输出数组里，此时左指针不动，右指针右移一位。
整个过程保持不变的性质是：区间 [0,left) 中的元素都不等于 val。当左右指针遍历完输入数组以后，left 的值就是输出数组的长度。

*/
func removeElement(nums []int, val int) int {
	left := 0
	// v 即 nums[right]
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

//优化
/*
如果左指针 left 指向的元素等于 val，此时将右指针 right 指向的元素复制到左指针 left 的位置，然后右指针 right 左移一位。
如果赋值过来的元素恰好也等于 val，可以继续把右指针 right 指向的元素的值赋值过来（左指针 left 指向的等于 val 的元素的位置继续被覆盖），
直到左指针指向的元素的值不等于 val 为止。

当左指针 left 和右指针 right 重合的时候，左右指针遍历完数组中所有的元素。
*/
func removeElement1(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
