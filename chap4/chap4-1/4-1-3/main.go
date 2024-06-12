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

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	sc.Buffer(buffer, maxN*12)

	xyr1 := scanLineInt()
	xyr2 := scanLineInt()

	floatXyr1 := make([]float64, len(xyr1))
	floatXyr2 := make([]float64, len(xyr2))
	for i := 0; i < 3; i++ {
		floatXyr1[i] = float64(xyr1[i])
		floatXyr2[i] = float64(xyr2[i])
	}

	dist := math.Sqrt(math.Pow(floatXyr1[0]-floatXyr2[0], 2) + math.Pow(floatXyr1[1]-floatXyr2[1], 2))

	if dist < abs(floatXyr1[2]-floatXyr2[2]) {
		fmt.Println(1)
	} else if dist == abs(floatXyr1[2]-floatXyr2[2]) {
		fmt.Println(2)
	} else if dist > abs(floatXyr1[2]-floatXyr2[2]) && dist < floatXyr1[2]+floatXyr2[2] {
		fmt.Println(3)
	} else if dist == floatXyr1[2]+floatXyr2[2] {
		fmt.Println(4)
	} else if dist > floatXyr1[2]+floatXyr2[2] {
		fmt.Println(5)
	}
}
