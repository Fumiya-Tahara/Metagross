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
	maxN   = 200000
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sc.Buffer(buffer, maxN*12)

	n := scanInt()
	a := scanLineInt()
	m := scanInt()

	cumSum := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			cumSum[i] = 0
			continue
		}
		cumSum[i] = cumSum[i-1] + a[i-1]
	}

	var s1 int
	var s2 int

	var ans int
	for i := 0; i < m; i++ {
		if i == 0 {
			s2 = scanInt()
			continue
		}
		s1 = s2
		s2 = scanInt()

		smaller := min(s1, s2)
		distance := abs(s1 - s2)

		ans += cumSum[smaller+distance-1] - cumSum[smaller-1]

		// for j := smaller - 1; j < smaller+distance-1; j++ {
		// 	ans += a[j]
		// }
	}

	fmt.Println(ans)
}
