package main

func main() {

}
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	maxStr := s[:1]
	for i := 0; i < len(s)-1; i++ {
		for j := i+1; j < len(s); j++ {
			if j-i+1 > len(maxStr) && isPalindrome(s[i:j+1]) {
				maxStr = s[i:j+1]
			}
		}
	}
	return maxStr
}
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[len(s)-i-1] != s[i] {
			return false
		}
	}
	return true
}