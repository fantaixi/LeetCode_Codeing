package main

/*
猜数字游戏的规则如下：

每轮游戏，我都会从1到n 随机选择一个数字。 请你猜选出的是哪个数字。
如果你猜错了，我会告诉你，你猜测的数字比我选出的数字是大了还是小了。
你可以通过调用一个预先定义好的接口 int guess(int num) 来获取猜测结果，返回值一共有 3 种可能的情况（-1，1或 0）：

-1：我选出的数字比你猜的数字小 pick < num
1：我选出的数字比你猜的数字大 pick > num
0：我选出的数字和你猜的数字一样。恭喜！你猜对了！pick == num
返回我选出的数字。
 */
func guessNumber(n int) int {
	//return sort.Search(n, func(i int) bool { return guess(i) <= 0 })
	left,right := 0,n
	for left <= right {
		mid := left + (right-left) >> 1
		if guess(mid) == 0 {
			return mid
		}else if guess(mid) == 1 {
			left = mid +1
		}else {
			right = mid-1
		}
	}
	return -1
}

func guess(i int) int {
	return 0
}