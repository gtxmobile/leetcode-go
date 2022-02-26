package main

func main() {
	
}
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	ans := n + 1
	if ans == 1 {
		return 0
	}
	start, end := 0, 0
	curs := 0
	for end < n {
		curs += nums[end]
		for curs >= s {
			tmp:= end-start+1
			if  tmp< ans {
				ans = tmp
			}
			curs -= nums[start]
			start++
		}
		end++
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}