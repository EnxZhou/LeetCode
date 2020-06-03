package problem0837

func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, K+W+1)
	for i := K; i <= N && i <= K+W-1; i++ {
		dp[i] = 1.0
	}
	if N < W+K-1 {
		dp[K-1] = float64(N-K+1) / float64(W)
	} else {
		dp[K-1] = 1.0
	}

	for i := K - 2; i >= 0; i-- {
		dp[i] = dp[i+1] + (dp[i+1]-dp[i+W+1])/float64(W)
	}
	return dp[0]
}
