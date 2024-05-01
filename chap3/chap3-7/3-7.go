package chap3_7

import "math"

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
