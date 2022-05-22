package main

func main() {

}
func peakIndexInMountainArray(arr []int) int {
	left,right := 1,len(arr)-1
	for left <= right {
		mid := (left + right) >> 1
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}else if arr[mid] < arr[mid-1] {
			right = mid - 1
		}else {
			left = mid +1
		}
	}
	return left
}