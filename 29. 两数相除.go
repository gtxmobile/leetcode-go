package main

func main() {

}
func divide(dividend int, divisor int) int {
	if dividend == -(1<<31) && divisor == -1 {
		return 1<<31 - 1
	}
	fu := -1
	dividend, divisor, fu = normal(dividend, divisor)
	ans := 0
	for dividend >= divisor {
		cnt := 0
		var tmp int
		for dividend >= divisor<<tmp{
			cnt = tmp
			tmp++
		}
		dividend -= divisor<<cnt
		ans += 1<<cnt
	}
	return ans*fu
}
func normal(dividend, divisor int) (int, int, int) {
	fu := -1
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		fu = 1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	return dividend, divisor, fu
}
