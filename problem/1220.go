package problem

func countVowelPermutation(n int) int {
	const mod int = 1e9 + 7
	vowel := []byte{'a', 'e', 'i', 'o', 'u'}
	dp := make([]map[byte]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make(map[byte]int, 5)
	}
	for i := 0; i < 5; i++ {
		dp[0][vowel[i]] = 0
		dp[1][vowel[i]] = 1
	}
	for i := 2; i <= n; i++ {
		dp[i]['a'] = (dp[i-1]['e'] + dp[i-1]['i'] + dp[i-1]['u']) % mod
		dp[i]['e'] = (dp[i-1]['a'] + dp[i-1]['i']) % mod
		dp[i]['i'] = (dp[i-1]['e'] + dp[i-1]['o']) % mod
		dp[i]['o'] = (dp[i-1]['i']) % mod
		dp[i]['u'] = (dp[i-1]['i'] + dp[i-1]['o']) % mod
	}
	ans := int64(0)
	for i := 0; i < 5; i++ {
		ans = (ans + int64(dp[n][vowel[i]])) % int64(mod)
	}
	return int(ans)
}
