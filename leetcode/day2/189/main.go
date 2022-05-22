package main

func main() {

}
func rotate(nums []int, k int)  {
	newArr := make([]int,len(nums))
	for i, v := range nums {
		newArr[(i+k)%len(nums)] = v
	}
	copy(nums,newArr)
}

func rotate1(nums []int, k int)  {
	k %= len(nums)
	reverse(nums)
	reverse(nums[k:])
	reverse(nums[:k])
}
func reverse(arr []int) {
	for i,n := 0,len(arr);i < n/2;i++ {
		arr[i],arr[n-i-1] = arr[n-i-1],arr[i]
	}
}