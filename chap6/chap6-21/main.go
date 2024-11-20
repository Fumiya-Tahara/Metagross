package main

import (
	"fmt"
	"math"
)

func main() {
	r := 2.0
	a := 1.0
	repeats := 5

	for i := 1; i <= repeats; i++ {
		x, y := a, math.Exp(a)

		slope := y
		intercept := y - slope*x

		nextA := (r - intercept) / slope
		fmt.Printf("Step #%d: a = %.15f -> %.15f\n", i, a, nextA)
		a = nextA
	}
}
