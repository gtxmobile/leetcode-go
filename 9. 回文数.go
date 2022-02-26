package main

func main() {
	println(isPalindrome(-121))
}
func isPalindrome(x int) bool {
	if x< 0{
		return false
	}
	ans := 0
	ori := x
	for x != 0{
		ans = ans * 10+x %10
		x /=10
	}
	return ans == ori
}