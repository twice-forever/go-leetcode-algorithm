package problem

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	t := make([]int, n)
	for i := 0; i < n; i++ {
		if t[i] == 1 {
			t[i] = 0
			break
		} else {
			t[i] = 1
		}
	}
	obstacleGrid[0] = t
	num := 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			num = 0
			obstacleGrid[i][0] = num
		} else {
			obstacleGrid[i][0] = num
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}
