package main

func main() {

}
func longestValidParentheses(s string) int {
	stack := []int{-1}
	ans := 0
	for i, b := range s {
		if b == ')' {
			stack = stack[:len(stack)-1]
			l := len(stack)
			if l > 0 {
				ans = max(ans, i-stack[l-1])
			} else {
				stack = append(stack, i)
			}
		} else {
			stack = append(stack, i)
		}
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
