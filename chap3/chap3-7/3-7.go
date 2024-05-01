package chap3_7

import (
	"math"
)

// カエルの足場問題
func Frog(numbers []int) int {
	dp := make([]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			dp[i] = 0
			continue
		}

		if i == 1 {
			dp[i] = dp[i-1] + int(math.Abs(float64(numbers[i]-numbers[i-1])))
			continue
		}

		patternA := dp[i-2] + int(math.Abs(float64(numbers[i]-numbers[i-2])))
		patternB := dp[i-1] + int(math.Abs(float64(numbers[i]-numbers[i-1])))

		if patternA < patternB {
			dp[i] = patternA
		} else {
			dp[i] = patternB
		}
	}

	return dp[len(numbers)-1]
}

func Stairs(n int) int {
	dp := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		if i == 0 {
			dp[i] = 0
			continue
		}
		if i == 1 {
			dp[i] = 1
			continue
		}
		if i == 2 {
			dp[i] = 2
			continue
		}

		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
