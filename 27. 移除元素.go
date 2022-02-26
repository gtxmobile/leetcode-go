package main

func main() {
	
}
func removeElement(nums []int, val int) int {
	l := len(nums)
	ans := 0
	for i := 0; i < l; i++ {
		if nums[i] != val{
			nums[ans] = nums[i]
			ans++
		}
	}
	return ans
}