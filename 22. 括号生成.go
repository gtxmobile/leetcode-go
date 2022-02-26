package main

import "fmt"

func main() {
	fmt.Print(generateParenthesis(3))
}
func generateParenthesis(n int) []string {
	var ans []string
	if n == 0 {
		return []string{""}
	}
	for i := 0; i < n; i++ {
		for _, sl := range generateParenthesis(i) {
			for _, sr := range generateParenthesis(n - i - 1) {
				ans = append(ans, fmt.Sprintf("(%s)%s", sl, sr))
			}
		}
	}
	return ans
}
