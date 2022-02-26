package main

func main() {
	
}
func maxArea(height []int) int {
	l,r := 0,len(height)-1
	ans := 0
	for l <r {
		if height[l]< height[r]{
			ans = max(ans, (r-l)*height[l])
			l++
		}else{
			ans = max(ans, (r-l)*height[r])
			r--
		}
	}
	return ans
}
func max(x,y int) int{
	if x > y {
		return x
	}
	return y
}