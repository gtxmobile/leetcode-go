package main

func main() {
	
}
func maxSubArray(nums []int) int {
	sum := 0
	ans := 0
	for _,v := range nums{
		if sum >0 {
			sum += v
		}else{
			sum = v
		}
		ans  = max(ans,sum)
	}
	return ans
}
func max(x,y int) int{
	if x> y{
		return x
	}
	return y
}