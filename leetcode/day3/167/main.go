package main

func main() {

}
func twoSum1(numbers []int, target int) []int {
	hashTable := map[int]int{}
	for i, val := range numbers {
		if p, ok := hashTable[target-val]; ok {
			return []int{p+1,i+1}
		}
		hashTable[val] = i
	}
	return nil
}
func twoSum(numbers []int, target int) []int {
	left,right := 0,len(numbers)-1
	for left < right {
		mid := (left + right) >> 1
		if numbers[left] + numbers[mid] > target {
			right = mid -1
		}else if numbers[mid]+numbers[right] < target {
			left = mid + 1
		}else if numbers[left] + numbers[right] > target {
			right--
		}else if numbers[left] + numbers[right] < target {
			left++
		}else {
			return []int{left+1,right+1}
		}
	}
	return []int{}
}