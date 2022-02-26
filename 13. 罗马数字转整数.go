package main

func main() {
	println(romanToInt("MCMXCIV"))
}
func romanToInt(s string) int {
	dict := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	i := 0
	l := len(s)
	ans := 0
	for i < l {
		if i+1 < l {
			v, found := dict[s[i:i+2]]
			if found {
				ans += v
				i += 2
				continue
			}
		}
		ans += dict[s[i:i+1]]
		i++
	}
	return ans
}

