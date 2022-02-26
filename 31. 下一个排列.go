package main

import (
	"fmt"
	"sort"
)

func main() {
	nextPermutation([]int{3,2,1})
}
func nextPermutation(nums []int)  {
	l := len(nums)
	i := l-1
	for ;i>=1;i--{
		if nums[i]>nums[i-1]{
			j := i
			for j <l && nums[j]>nums[i-1]{
				j++
			}
			nums[i-1], nums[j-1] = nums[j-1],nums[i-1]
			break
		}
	}
	sort.Ints(nums[i:])
	fmt.Print(nums)
}
