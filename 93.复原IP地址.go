package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(restoreIpAddresses("25525511135"))
}
func restoreIpAddresses(s string) []string {
	var ans []string
	lenth := len(s)
	if lenth > 12 || lenth < 4 {
		return ans
	}
	for i := 0; i < 3; i++ {
		for j := i + 1; j < lenth-2; j++ {
			for k := j + 1; k < lenth-1; k++ {
				if k < lenth-4 {
					continue
				}
				if check(&s, i, j, k, lenth) {
					res := fmt.Sprintf("%s.%s.%s.%s", s[0:i+1], s[i+1:j+1], s[j+1:k+1], s[k+1:lenth])
					//println(res)
					ans = append(ans, res)
				}
			}
		}
	}
	return ans
}
func check(s *string, i, j, k, l int) bool {
	if !valid(s, 0, i+1) {
		return false
	}
	if !valid(s, i+1, j+1) {
		return false
	}
	if !valid(s, j+1, k+1) {
		return false
	}
	if !valid(s, k+1, l) {
		return false
	}
	return true
}
func valid(s *string, b, e int) bool {
	v1, _ := strconv.Atoi((*s)[b:e])
	if v1 < 0 || v1 > 255 {
		return false
	}
	if (*s)[b] == '0' && e-b > 1 {
		return false
	}
	return true
}
