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
	maxN   = 1000000
	maxVal = 1000000
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

func scanInt64() int64 {
	sc.Scan()
	i, err := strconv.ParseInt(sc.Text(), 10, 64)
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

func scanLineInt64() []int64 {
	s := scanStr()
	strList := strings.Split(s, " ")

	slice := make([]int64, len(strList))
	for i, v := range strList {
		num, _ := strconv.ParseInt(string(v), 10, 64)
		slice[i] = num
	}

	return slice
}

func main() {
	sc.Buffer(buffer, maxVal)

	HW := scanLineInt()
	H := HW[0]
	W := HW[1]

	A := make([][]int, H+1)
	calcRow := make([]int, H+1)
	calcColumn := make([]int, W+1)
	for i := 1; i <= H; i++ {
		inputA := scanLineInt()
		A[i] = make([]int, W+1)

		for j := 1; j <= W; j++ {
			A[i][j] = inputA[j-1]
			calcRow[i] += A[i][j]
			calcColumn[j] += A[i][j]
		}
	}

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			calculated := calcRow[i] + calcColumn[j] - A[i][j]
			fmt.Printf("%v ", calculated)
		}
		fmt.Printf("\n")
	}
}
