package main

func main() {
}
func findCircleNum(grid [][]int) int {
	ans := 0
	row := len(grid)
	if row == 0 {
		return row
	}
	visit := make([][]int, row)
	for i := 0; i < row; i++ {
		visit[i] = make([]int, row)
	}
	for i := 0; i < row; i++ {
		tmp := bfs(i, row, &visit, &grid)
		if tmp >= 1 {
			ans++
		}
	}
	return ans
}
func bfs(x, row int, vis, grid *[][]int) int {
	ret := 0
	for y := 0; y < row; y++ {
		if (*vis)[x][y] == 1 {
			continue
		}
		(*vis)[x][y] = 1
		if (*grid)[x][y] == 1 {
			ret = ret + 1 + bfs(y, row, vis, grid)
		}
	}
	return ret
}
