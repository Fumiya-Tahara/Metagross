package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	cp := make([][2]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&cp[i][0], &cp[i][1])
	}

	var Q int
	fmt.Scan(&Q)

	lr := make([][2]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Scan(&lr[i][0], &lr[i][1])
	}

	sum1 := make([]int, N+1)
	sum2 := make([]int, N+1)
	for i, v := range cp {
		if v[0] == 1 {
			sum1[i+1] = sum1[i] + v[1]
			sum2[i+1] = sum2[i]
		} else if v[0] == 2 {
			sum1[i+1] = sum1[i]
			sum2[i+1] = sum2[i] + v[1]
		}
	}

	for i := 0; i < Q; i++ {
		fmt.Println(sum1[lr[i][1]]-sum1[lr[i][0]-1], sum2[lr[i][1]]-sum2[lr[i][0]-1])
	}
}
