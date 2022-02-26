package main

import "strings"

func main() {
	println(multiply("123", "456"))
}
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	ans := make([]byte, l1+l2)
	var c byte = 0
	for i := l1 - 1; i >= 0; i-- {
		intn1 := num1[i] - '0'
		for j := l2 - 1; j >= 0; j-- {
			ji := intn1*(num2[j]-'0') + c + ans[i+j+1]
			c = ji / 10
			ans[i+j+1] = ji % 10
		}
		ans[i] = c
		c = 0
		//output(ans)
	}
	for i, v := range ans {
		ans[i] = v + '0'
	}
	return strings.TrimLeft(string(ans), "0")
}
func output(a []byte) {
	for _, v := range a {
		print(v)
	}
	println()
}
