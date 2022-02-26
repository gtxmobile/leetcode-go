package main

import "strings"

func main() {
	println(intToRoman(1994))
}
func intToRoman(num int) string {
	zimu := []string{
		"M",
		"CM",
		"D",
		"CD",
		"C",
		"XC",
		"L",
		"XL",
		"X",
		"IX",
		"V",
		"IV",
		"I",
	}
	shuzi := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}
	var ans []string
	for i, v := range shuzi {
		num = count(num, v, &ans, zimu[i])
	}
	return strings.Join(ans, "")
}

func count(num, cnt int, ans *[]string, luo string) int {
	for ; num >= cnt; num -= cnt {
		*ans = append(*ans, luo)
	}
	return num
}
