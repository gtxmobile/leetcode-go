package main
func twoSum(nums []int, target int) []int {
	d := make(map[int]int)
	for i,v := range(nums){
		d[v] = i
	}
	for j,w := range(nums){
		diff := target-w
		index,found := d[diff]
		if found && index>j{
			return []int{j,index}
		}
	}
	return []int{}
}