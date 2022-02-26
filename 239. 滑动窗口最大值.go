package main

import "fmt"

func main() {
	fmt.Print(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
}
func maxSlidingWindow(nums []int, k int) []int {
	var ans []int
	var stack []int
	for i,v := range nums{
		for len(stack) > 0 && v > nums[stack[len(stack)-1]]{
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if stack[0] == i-k{
			stack = stack[1:]
		}
		if i>=k-1{
			ans = append(ans, nums[stack[0]])
		}
	}
	return ans
}
