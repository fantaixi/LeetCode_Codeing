package main

/*
给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母target，请你寻找在这一有序列表里比目标字母大的最小字母。
在比较时，字母是依序循环出现的。举个例子：
如果目标字母 target = 'z' 并且字符列表为letters = ['a', 'b']，则答案返回'a'
*/
//暴力
func nextGreatestLetter(letters []byte, target byte) byte {
	for _, b := range letters {
		if b > target {
			return b
		}
	}
	return letters[0]
}

//二分
// 找比目标数最小的数字， 不需要与右边字母比较
func nextGreatestLetter1(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	ans := target
	for left <= right {
		mid := left + (right-left)>>1
		if letters[mid] > target { //向左走，找是否还有更小的
			ans = letters[mid] //此数字是满足条件的但需要继续查找是否有更好的数字
			right = mid - 1
		} else { //当mid≤target时候直接排除掉mid
			left = mid + 1
		}
	}
	// 还需要判断最后跳出的字母是否符合条件
	if ans == target { // 说明整个数字都没有比target大的
		ans = letters[0]
	}
	return ans
}
