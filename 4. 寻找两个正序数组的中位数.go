package main

func main() {

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 1 {
		return float64(getKthElement(nums1, nums2, l/2+1))
	} else {
		return (float64(getKthElement(nums1, nums2, l/2)) + float64(getKthElement(nums1, nums2, l/2+1))) / 2.0
	}
}
func getKthElement(nums1, nums2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		if idx1 == len(nums1) {
			return nums2[idx2+k-1]
		}
		if idx2 == len(nums2) {
			return nums1[idx1+k-1]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}
		mid := k / 2
		new_idx1 := min(idx1+mid, len(nums1)) - 1
		new_idx2 := min(idx2+mid, len(nums2)) - 1
		if nums1[new_idx1] < nums2[new_idx2] {
			k -= new_idx1 - idx1 + 1
			idx1 = new_idx1 + 1
		} else {
			k -= new_idx2 - idx2 + 1
			idx2 = new_idx2 + 1
		}
	}
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}