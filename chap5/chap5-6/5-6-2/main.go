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

func scanFloat64() float64 {
	s := scanStr()
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func scanLineFloat64() []float64 {
	s := scanStr()
	strList := strings.Split(s, " ")

	slice := make([]float64, len(strList))
	for i, v := range strList {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatal(err)
		}
		slice[i] = num
	}

	return slice
}

func main() {
	sc.Buffer(buffer, maxVal)

	mod := int64(1000000007)

	N := scanInt64()
	A := scanLineInt64()

	multiplier := int64(1)
	var answer int64
	for i := int64(0); i < N; i++ {
		multiplier %= mod
		answer += multiplier * A[i]
		answer %= mod
		multiplier *= 2
	}

	fmt.Println(answer)
}
