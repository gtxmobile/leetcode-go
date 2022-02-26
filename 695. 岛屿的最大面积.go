package main

func main() {

}
func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	row := len(grid)
	if row == 0 {
		return row
	}
	col := len(grid[0])
	visit := make([][]int, row)
	for i := 0; i < row; i++ {
		visit[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			tmp := bfs(i, j, row, col, &visit, &grid)
			if tmp > ans {
				ans = tmp
			}
		}
	}
	return ans
}
func bfs(x, y, row, col int, vis, grid *[][]int) int {
	if x < 0 || x >= row || y < 0 || y >= col {
		return 0
	}
	if (*vis)[x][y] == 1 {
		return 0
	}
	(*vis)[x][y] = 1
	if (*grid)[x][y] == 0 {
		return 0
	}
	return 1 + bfs(x-1, y, row, col, vis, grid) + bfs(x+1, y, row, col, vis, grid) + bfs(x, y-1, row, col, vis, grid) + bfs(x, y+1, row, col, vis, grid)
}
