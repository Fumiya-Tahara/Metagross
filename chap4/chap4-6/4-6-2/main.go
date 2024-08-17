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

func modPow(a, b, m int) int {
	p := a
	answer := 1

	for i := 0; i < 30; i++ {
		if (b & (1 << i)) != 0 {
			answer *= p
			answer %= m
		}

		p *= p
		p %= m
	}

	return answer
}

func division(a, b, m int) int {
	return (a * modPow(b, m-2, m)) % m
}

func factorialMod(n, mod int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		result %= mod
	}

	return result
}

func main() {
	sc.Buffer(buffer, maxVal)

	mod := 1000000007

	XY := scanLineInt()
	X := XY[0]
	Y := XY[1]

	a := (2*X - Y) / 3
	b := (2*Y - X) / 3

	if ((2*X-Y)%3 != 0 || a < 0) || ((2*Y-X)%3 != 0 || b < 0) {
		fmt.Println(0)

		return
	}

	deno := factorialMod(a, mod) * factorialMod(b, mod)
	deno %= mod
	mole := factorialMod(a+b, mod)

	answer := division(mole, deno, mod)

	fmt.Println(answer)
}
