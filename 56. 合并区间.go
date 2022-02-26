package main

import "sort"

func main() {
	
}
func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//fmt.Print(intervals)
	ans := [][]int{}
	for _, v := range intervals {
		l := len(ans)
		if l == 0 {
			ans = append(ans, v)
			continue
		}
		if ans[l-1][1]< v[0]{
			ans = append(ans,v)
		}else if ans[l-1][1]< v[1]{
			ans[l-1][1] = v[1]
		}

	}
	return ans
}