package main

func main() {

}
func reverseString(s []byte)  {
	for left,rigth := 0,len(s)-1; left<rigth  ; left,rigth = left+1,rigth-1 {
		s[left],s[rigth] = s[rigth],s[left]
	}
}
