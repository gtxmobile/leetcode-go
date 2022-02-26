package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(fourSum([]int{1, 0, -1, 0, -2, 2},0))
}
func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	lenth := len(nums)
	if lenth < 4 {
		return ans
	}
	sort.Ints(nums)
	left, right := 0, 0
	last := nums[0] - 1
	for j := 0; j < lenth-3; j++ {
		sum := target-nums[j]
		if nums[j] != last {
			last2 := nums[j+1] - 1
			for i := j + 1; i < lenth-2; i++ {
				if nums[i] != last2 {
					left = i + 1
					right = lenth - 1
					tmp2 := sum - nums[i]
					for left < right {
						tmp := nums[left] + nums[right]
						if tmp == tmp2 {
							ans = append(ans, []int{nums[j], nums[i], nums[left], nums[right]})
							for left < right {
								left++
								if nums[left] != nums[left-1] {
									break
								}
							}
							for left < right {
								right--
								if nums[right] != nums[right+1] {
									break
								}
							}
						} else if tmp < tmp2 {
							left++
						} else {
							right--
						}
					}
					last2 = nums[i]
				}
			}
			last = nums[j]
		}
	}
	return ans
}
