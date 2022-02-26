package main

import "sort"

func main() {

}
func threeSumClosest(nums []int, target int) int {
	var ans = 1000000
	lenth := len(nums)
	sort.Ints(nums)
	left, right := 0, 0
	ori := target
	for i := 0; i < lenth-2; i++ {
		left = i + 1
		right = lenth - 1
		target := ori - nums[i]
		for left < right {
			tmp := nums[left] + nums[right]
			if tmp == target {
				return ori
			} else {
				if abs(ans-ori)>abs(tmp-target){
					ans = tmp + nums[i]
				}
				if tmp < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return ans
}
func abs(x int) int{
	if x < 0{
		return -x
	}
	return x
}