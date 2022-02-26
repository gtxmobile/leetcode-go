package main

func main() {
	
}
func canJump(nums []int) bool {
	max_pos :=0
	l := len(nums)
	for i, n := range nums{
		if max_pos >=i{
			if i+n > max_pos{
				max_pos = i+n
			}
			if max_pos >= l-1{
				return true
			}
		}
	}
	return false
}