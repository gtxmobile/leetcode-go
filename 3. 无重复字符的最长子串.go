package main

func main() {
	println(lengthOfLongestSubstring("tmmzuxt"))

}
func lengthOfLongestSubstring(s string) int {
	maxret := 0
	left := 0
	lastSeen := make(map[int32]int)
	for i, v := range s {
		last, found := lastSeen[v]
		if found && last >= left {
			left = last + 1
		} else {
			if i-left+1 > maxret {
				maxret = i - left + 1
			}
		}
		lastSeen[v] = i
	}
	return maxret
}
