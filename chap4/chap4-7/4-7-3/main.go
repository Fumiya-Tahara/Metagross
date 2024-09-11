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

type Matrix struct {
	p [3][3]int
}

func Multiplication(A, B Matrix) Matrix {
	C := Matrix{}

	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			for j := 0; j < 3; j++ {
				C.p[i][j] += A.p[i][k] * B.p[k][j]
				C.p[i][j] %= 1000000007
			}
		}
	}

	return C
}

func Power(A Matrix, n int) Matrix {
	P := A

	Q := Matrix{}
	flag := false

	for i := 0; i < 60; i++ {
		if (n & (1 << i)) != 0 {
			if !flag {
				Q = P
				flag = true
			} else {
				Q = Multiplication(Q, P)
			}
		}
		P = Multiplication(P, P)
	}

	return Q
}

func main() {
	sc.Buffer(buffer, maxVal)

	N := scanInt()

	A := Matrix{}
	A.p[0][0] = 1
	A.p[0][1] = 1
	A.p[0][2] = 1
	A.p[1][0] = 1
	A.p[2][1] = 1

	B := Power(A, N-3)

	answer := (B.p[0][0]*2 + B.p[0][1] + B.p[0][2]) % 1000000007

	fmt.Println(answer)
}
