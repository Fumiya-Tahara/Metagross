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
	maxN   = 10000000
	maxVal = 10000000
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

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) push(n int) {
	q.data = append(q.data, n)
}

func (q *Queue) pop() int {
	if len(q.data) == 0 {
		return -1
	}
	n := q.data[0]
	q.data = q.data[1:]
	return n
}

func main() {
	sc.Buffer(buffer, maxVal)

	maxAge := 120

	NM := scanLineInt()
	N := NM[0]
	M := NM[1]

	G := make([][]int, N+1)

	for i := 1; i <= M; i++ {
		ab := scanLineInt()
		a := ab[0]
		b := ab[1]

		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	dist := make([]int, N+1)
	for i := 0; i <= N; i++ {
		dist[i] = -1
	}

	Q := NewQueue()
	Q.push(1)
	dist[1] = 0

	for !Q.isEmpty() {
		pos := Q.pop()
		for i := 0; i < len(G[pos]); i++ {
			next := G[pos][i]
			if dist[next] == -1 {
				dist[next] = dist[pos] + 1
				Q.push(next)
			}
		}
	}

	for i := 1; i <= N; i++ {
		if dist[i] == -1 {
			fmt.Println(maxAge)
		} else {
			if dist[i] > maxAge {
				fmt.Println(maxAge)
			} else {
				fmt.Println(dist[i])
			}
		}
	}

}
