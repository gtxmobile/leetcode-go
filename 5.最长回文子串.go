package main

func main() {
	
}
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start,end := 0,0
	for i, _ := range s {
		l1, r1 := expand(s, i, i)
		l2, r2 := expand(s, i, i+1)
		if r1-l1>end-start{
			start = l1
			end = r1
		}
		if r2-l2>end-start{
			start = l2
			end = r2
		}
	}
	return s[start:end+1]
}
func expand(s string, left, right int) (int, int) {
	for left>=0 && right< len(s) && s[left]==s[right]{
		left--
		right++
	}
	return left +1,right-1
}