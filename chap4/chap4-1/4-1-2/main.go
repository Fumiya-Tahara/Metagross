package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	maxN   = 2000
	maxVal = 1000000000
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	buffer = make([]byte, maxN*12)
)

func scanStr() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func scanLineStr() []string {
	s := scanStr()

	return strings.Split(s, " ")
}

func scanLineInt() []int {
	s := scanStr()
	strList := strings.Split(s, " ")

	slice := make([]int, len(strList))
	for i, v := range strList {
		num, _ := strconv.Atoi(string(v))
		slice[i] = num
	}

	return slice
}

func main() {
	sc.Buffer(buffer, maxN*12)

	n := scanInt()

	xyList := make([][]int, n)
	for i := 0; i < n; i++ {
		xy := scanLineInt()
		xyList[i] = xy
	}

	min := maxVal
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x1 := xyList[i][0]
			y1 := xyList[i][1]
			x2 := xyList[j][0]
			y2 := xyList[j][1]

			xDiff := x2 - x1
			yDiff := y2 - y1

			distance := xDiff*xDiff + yDiff*yDiff
			if distance < min {
				min = distance
			}
		}
	}

	floatDist := float64(min)
	result := math.Sqrt(floatDist)
	fmt.Println(result)
}
