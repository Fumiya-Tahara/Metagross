package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	var n int
	fmt.Scan(&n)

	t := make([]int, n+1)
	var tSum int
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i+1])
		tSum += t[i+1]
	}

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, tSum+1)
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= tSum; j++ {
			if j < t[i] {
				if dp[i-1][j] {
					dp[i][j] = true
				}
			} else if j >= t[i] {
				if dp[i-1][j] || dp[i-1][j-t[i]] {
					dp[i][j] = true
				}
			}
		}
	}

	answer := tSum
	for i := 0; i <= tSum; i++ {
		if dp[n][i] {
			cookingTime := max(i, tSum-i)
			answer = min(answer, cookingTime)
		}
	}

	fmt.Println(answer)
}
