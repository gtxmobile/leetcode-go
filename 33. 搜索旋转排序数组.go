package main

func main() {

}
func search(nums []int, target int) int {
	ret := -1
	lenth := len(nums)
	if lenth == 0 {
		return ret
	}
	l, r := 0, lenth-1
	for l <= r {
		mid := (l + r) / 2
		mid_v := nums[mid]
		if mid_v == target {
			return mid
		}
		if nums[l] <= mid_v {
			if nums[l] <= target && target < mid_v {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if mid_v < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return ret
}
