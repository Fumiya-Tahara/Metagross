package chap3_7

import (
	"fmt"
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

// 階段の登り方問題
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

// ナップザック問題
func Knapsack(n, w int, items [][]int) int {
	dp := make([][]int, n+1)
	dp[0] = make([]int, w+1)

	for i := 1; i <= n; i++ {
		dp[i] = make([]int, w+1)
		for j := 1; j <= w; j++ {
			if j < items[i-1][0] {
				dp[i][j] = dp[i-1][j]
				continue
			}
			isSelectedVal := dp[i-1][j-items[i-1][0]] + items[i-1][1]
			isNotSelectedVal := dp[i-1][j]
			if isSelectedVal > isNotSelectedVal {
				dp[i][j] = isSelectedVal
			} else {
				dp[i][j] = isNotSelectedVal
			}
		}
	}

	return dp[n][w]
}

// 3.7.4
func PartialSum(N, S int, A []int) {
	dp := make([][]bool, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]bool, S+1)
	}
	dp[0][0] = true

	for i := 1; i < N+1; i++ {
		for j := 0; j < S+1; j++ {
			if j < A[i-1] {
				dp[i][j] = dp[i-1][j]
				continue
			}
			if dp[i-1][j] || dp[i-1][j-A[i-1]] {
				dp[i][j] = true
				continue
			}
			dp[i][j] = false
		}
	}

	if dp[N][S] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 3.7.6
func Taro(N int, A []int) int {
	dp := make([]int, N+1)
	dp[0] = 0
	dp[1] = A[0]

	for i := 2; i < N+1; i++ {
		suruhi := dp[i-2] + A[i-1]
		shinaihi := dp[i-1]

		if suruhi > shinaihi {
			dp[i] = suruhi
		} else {
			dp[i] = shinaihi
		}
	}

	return dp[N]
}
