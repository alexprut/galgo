package graph

import (
	"galgo/algorithms/utils"
	"math"
)

func FloydWarshall(adj [][]int) [][]int {
	n := len(adj)
	dp := adj
	for i := 0; i < n; i++ {
		dp[i][i] = 0
	}

	for k := 0; k < n; k++ {
		tmp := make([][]int, n, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				var tmpVal int
				if utils.MaxInt(dp[i][k], dp[k][j]) == math.MaxInt64 {
					tmpVal = math.MaxInt64
				} else {
					tmpVal = dp[i][k] + dp[k][j]
				}
				tmp[i][j] = utils.MinInt(dp[i][j], tmpVal)
			}
		}
		dp = tmp
	}

	return dp
}
