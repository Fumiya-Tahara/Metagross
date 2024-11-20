package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	B := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		fmt.Scan(&B[i])
	}

	A := make([]int, n)
	var sum int
	for i := 0; i < n; i++ {
		if i == 0 {
			A[i] = B[i]
			sum += B[i]
			continue
		}

		if i == n-1 {
			A[i] = B[i-1]
			sum += B[i-1]
			break
		}

		A[i] = min(B[i-1], B[i])
		sum += min(B[i-1], B[i])
	}

	fmt.Println(sum)
}
