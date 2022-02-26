package main

func main() {

}
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	x,y := row-1,0
	for x >=0 && y<col{
		if matrix[x][y]==target{
			return true
		}else if matrix[x][y]>target{
			x--
		}else{
			y++
		}
	}
	return false
}
