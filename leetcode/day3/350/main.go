package main

func main() {

}
func intersect(nums1 []int, nums2 []int) []int {
	numCnt := make(map[int]int)
	res := []int{}
	for i := range nums1 {
		numCnt[nums1[i]]++
	}
	for j := range nums2 {
		if count, ok := numCnt[nums2[j]]; ok && count > 0 {
			res = append(res,nums2[j])
			numCnt[nums2[j]]--
		}
	}
	return res
}