package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	xy := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&xy[i][0], &xy[i][1])
	}

	ab := [2]int{}
	fmt.Scan(&ab[0], &ab[1])

	var count int
	for i := 0; i < n; i++ {
		var xBottom, yBottom, xTop, yTop int

		if i == n-1 {
			xBottom = xy[i][0] - ab[0]
			yBottom = xy[i][1] - ab[1]
			xTop = xy[0][0] - ab[0]
			yTop = xy[0][1] - ab[1]
		} else {
			xBottom = xy[i][0] - ab[0]
			yBottom = xy[i][1] - ab[1]
			xTop = xy[i+1][0] - ab[0]
			yTop = xy[i+1][1] - ab[1]
		}

		if yBottom > yTop {
			xBottom, xTop = xTop, xBottom
			yBottom, yTop = yTop, yBottom
		}

		if (yBottom <= 0 && 0 < yTop) && xBottom*yTop-xTop*yBottom < 0 {
			count++
		}
	}

	if count%2 == 1 {
		fmt.Println("INSIDE")
		return
	}

	fmt.Println("OUTSIDE")
}
