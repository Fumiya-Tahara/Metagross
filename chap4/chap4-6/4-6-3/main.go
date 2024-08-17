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

func modPow(a, b, m int64) int64 {
	var p int64 = a
	var answer int64 = 1

	for i := 0; i < 60; i++ {
		if (b & (1 << i)) != 0 {
			answer *= p
			answer %= m
		}

		p *= p
		p %= m
	}

	return answer
}

func division(a, b, m int64) int64 {
	return (a * modPow(b, m-2, m)) % m
}

func main() {
	sc.Buffer(buffer, maxVal)

	var mod int64 = 1000000007

	n := scanInt64()

	var dino int64 = 3
	mole := modPow(4, n+1, mod) - 1

	answer := division(mole, dino, mod)

	fmt.Println(answer)
}
