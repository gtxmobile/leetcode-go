package main

func main() {

}
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 || l == 1 {
		return l
	}
	pre := nums[0]
	ans := 1
	for i := 1; i < l; i++ {
		if nums[i] != pre{
			pre = nums[i]
			nums[ans] = nums[i]
			ans++
		}
	}
	return ans
}
