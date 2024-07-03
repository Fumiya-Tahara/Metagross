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

func Newton() {
	var r float64 = 2.0
	var a float64 = 2.0

	for i := 0; i < 5; i++ {
		var zahyou_x float64 = a
		var zahyou_y float64 = a * a

		var sessen_a float64 = 2.0 * zahyou_x
		var sessen_b float64 = zahyou_y - sessen_a*zahyou_x

		var next_a float64 = (r - sessen_b) / sessen_a
		fmt.Printf("Step #%d: a = %.12f -> %.12f\n", i+1, a, next_a)
		a = next_a
	}

	fmt.Println(a)
}

func BinarySearch() {
	var l float64 = 1.0
	var r float64 = 2.0

	var m float64
	for i := 0; i < 20; i++ {
		m = (l + r) / 2
		if m*m < 2 {
			l = m
		} else {
			r = m
		}
	}

	fmt.Println(m)
}

func main() {
	Newton()
	fmt.Println("**********************")
	BinarySearch()

}
