package p9

func MinCost1(s1, s2 string, ic, dc, rc int) int {
	m, n := len(s1), len(s2)
	dp := initdp(m+1, n+1)
	for i := 1; i <= m; i++ {
		dp[i][0] = i * dc
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j * ic
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i][j-1]+ic,
					dp[i-1][j]+dc,
					dp[i-1][j-1]+rc)
			}
		}
	}
	return dp[m][n]
}

func initdp(m, n int) [][]int {
	base := make([]int, m*n)
	dp := make([][]int, m)
	for i := range dp {
		dp[i], base = base[0:n], base[n:]
	}
	return dp
}

func min(a ...int) int {
	m := a[0]
	for _, v := range a[1:] {
		if v < m {
			m = v
		}
	}
	return m
}
