package problem

func minimumTotal(triangle [][]int) int {
	ans := triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] += triangle[i-1][0]
		triangle[i][len(triangle[i])-1] += triangle[i-1][len(triangle[i-1])-1]
		ans = triangle[i][0]
		for j := 1; j < len(triangle[i]); j++ {
			if j+1 < len(triangle[i]) {
				triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
			}
			if i+1 == len(triangle) && ans > triangle[i][j] {
				ans = triangle[i][j]
			}
		}
	}
	return ans
}
