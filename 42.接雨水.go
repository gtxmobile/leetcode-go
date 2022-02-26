package main

import "math"

func main() {
	println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
func trap(height []int) int {
	ans, max_l, max_r := 0, 0, 0
	length := len(height) - 1
	r_arr := make([]int, length+1)
	for i := length; i >= 0; i-- {
		if height[i] > max_r {
			max_r = height[i]
		}
		r_arr[i] = max_r
	}
	for i, v := range height {
		if v > max_l {
			max_l = v
		}
		cnt := int(math.Min(float64(max_l), float64(r_arr[i])))
		ans = ans + cnt - v
	}
	return ans
}
