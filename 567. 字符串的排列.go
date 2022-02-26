package main

func main() {
	println(checkInclusion("abcdxabcde", "abcdeabcdx"))
}
func checkInclusion(s1 string, s2 string) bool {
	l1 := len(s1)
	l, r := 0, 0
	visit := [26]int{}
	for _, s := range s1 {
		visit[s-'a'] += 1
	}
	for _, s := range s2 {
		c := s - 'a'
		visit[c] -= 1
		r++
		for l < r && visit[c] < 0 {
			visit[s2[l]-'a'] += 1
			l++
		}
		if r-l-l1 == 0 {
			return true
		}
	}
	return false
}
