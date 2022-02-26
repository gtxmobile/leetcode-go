package main

func main() {
	
}
func maximalSquare(matrix [][]byte) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])
	if col ==0 {
		return 0
	}
	dp := make([][]byte, row)
	var maxBian byte = 0
	for i := 0; i < row; i++ {
		dp[i] = make([]byte, col)
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(min(dp[i-1][j], dp[i-1][j-1]), min(dp[i-1][j-1], dp[i][j-1])) + 1
				}
				maxBian = max(maxBian, dp[i][j])
			}
		}
	}
	return int(maxBian * maxBian)
}
func min(x, y byte) byte {
	if x < y {
		return x
	}
	return y
}
func max(x, y byte) byte {
	if x > y {
		return x
	}
	return y
}