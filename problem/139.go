package problem

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if wordDictSet[s[j:i]] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
