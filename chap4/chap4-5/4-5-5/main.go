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
	maxN   = 100000
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
	sc.Buffer(buffer, maxVal)

	nm := scanLineInt()
	n := nm[0]
	m := nm[1]

	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		ab := scanLineInt()
		a[i] = ab[0]
		b[i] = ab[1]
	}

	g := make([][]int, n+1)
	for i := 0; i < m; i++ {
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	var answer int
	for i := 1; i <= n; i++ {
		var count int
		for _, v := range g[i] {
			if v < i {
				count++
			}
		}
		if count == 1 {
			answer++
		}
	}

	fmt.Println(answer)
}
