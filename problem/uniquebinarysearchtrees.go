package problem

func NumTrees(n int) int {
	G := make([]int, n+1)

	G[0], G[1] = 1, 1
	for j := 2; j <= n; j++ {
		for i := 1; i <= j; i++ {
			G[j] += G[i-1] * G[j-i]
		}
	}
	return G[n]
}
