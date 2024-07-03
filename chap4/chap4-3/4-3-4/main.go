package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	maxN   = 500000
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

func cubeRoot(n float64) float64 {
	var r float64 = n
	var a float64 = 2.0

	for i := 0; i < 100; i++ {
		var zahyou_x float64 = a
		var zahyou_y float64 = a * a * a

		var sessen_a float64 = 3.0 * zahyou_x * zahyou_x
		var sessen_b float64 = zahyou_y - sessen_a*zahyou_x

		var next_a float64 = (r - sessen_b) / sessen_a
		a = next_a
	}

	return a
}

func main() {
	var r float64 = 10.0
	var a float64 = 2.0

	for i := 0; i < 5; i++ {
		var zahyou_x float64 = a
		var zahyou_y float64 = a * a * a * cubeRoot(a)

		var sessen_a float64 = 10 / 3 * zahyou_x * zahyou_x * cubeRoot(a)
		var sessen_b float64 = zahyou_y - sessen_a*zahyou_x

		var next_a float64 = (r - sessen_b) / sessen_a
		a = next_a
	}

	fmt.Println(a)
}
