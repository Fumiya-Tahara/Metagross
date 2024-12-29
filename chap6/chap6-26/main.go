package main

import (
	"fmt"
)

func main() {
	var Q int
	fmt.Scan(&Q)

	for query := 1; query <= Q; query++ {
		var X, Y, Z float64
		var T int
		fmt.Scan(&X, &Y, &Z, &T)

		ans := [3][3]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

		for T > 0 {
			num := 1
			A := [3][3]float64{{1 - X, Y, 0}, {0, 1 - Y, Z}, {X, 0, 1 - Z}}

			for num*2 <= T {
				tmp := A
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						A[i][j] = tmp[i][0]*tmp[0][j] + tmp[i][1]*tmp[1][j] + tmp[i][2]*tmp[2][j]
					}
				}
				num *= 2
			}

			T -= num

			tmp2 := ans
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					ans[i][j] = tmp2[i][0]*A[0][j] + tmp2[i][1]*A[1][j] + tmp2[i][2]*A[2][j]
				}
			}
		}

		fmt.Printf("%.10f %.10f %.10f\n", ans[0][0]+ans[0][1]+ans[0][2], ans[1][0]+ans[1][1]+ans[1][2], ans[2][0]+ans[2][1]+ans[2][2])
	}
}
