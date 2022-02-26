package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
func threeSum(nums []int) [][]int {
	var ans [][]int
	lenth := len(nums)
	if lenth < 3 {
		return ans
	}
	sort.Ints(nums)
	left, right := 0, 0
	last := nums[0] - 1
	for i := 0; i < lenth-2; i++ {
		if nums[i] != last {
			left = i + 1
			right = lenth - 1
			target := -nums[i]
			for left < right {
				tmp := nums[left] + nums[right]
				if tmp == target {
					ans = append(ans, []int{nums[i], nums[left], nums[right]})
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
				} else if tmp < target {
					left++
				} else {
					right--
				}
			}
			last = nums[i]
		}
	}
	return ans
}
