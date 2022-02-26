package main

func main() {
	
}
func strStr(haystack string, needle string) int {
	ans := -1
	lh := len(haystack)
	ln := len(needle)
	if lh< ln{
		return ans
	}
	for i := 0;i<=lh-ln;i++{
		if haystack[i:ln+i] == needle{
			return i
		}
	}
	return ans

}