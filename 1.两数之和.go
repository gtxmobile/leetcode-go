package main

import "fmt"

func main(){
	fmt.Print(twoSum([]int{2,2},4))
}
func twoSum(nums []int, target int) []int {
	d := make(map[int]int)
	for j, w := range nums {
		index, found := d[target - w]
		if found  {
			return []int{j, index}
		}
		d[w] = j
	}
	return []int{}
}
