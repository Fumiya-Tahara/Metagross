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

type Queue struct {
	data [][]int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) push(yx []int) {
	q.data = append(q.data, yx)
}

func (q *Queue) pop() []int {
	if len(q.data) == 0 {
		return nil
	}
	yx := q.data[0]
	q.data = q.data[1:]
	return yx
}

func main() {
	sc.Buffer(buffer, maxVal)

	RC := scanLineInt()
	R := RC[0]
	C := RC[1]

	s := scanLineInt()
	g := scanLineInt()

	c := make([][]string, R+1)
	for i := 1; i <= R; i++ {
		char := scanStr()

		chars := make([]string, C+1)
		for j := 1; j <= C; j++ {
			chars[j] = string(char[j-1])
		}

		c[i] = chars
	}

	dist := make([][]int, R+1)
	for i := 1; i <= R; i++ {
		dist[i] = make([]int, C+1)
		for j := 1; j <= C; j++ {
			dist[i][j] = -1
		}
	}

	q := NewQueue()
	q.push(s)
	dist[s[0]][s[1]] = 0

	for !q.isEmpty() {
		yx := q.pop()
		y := yx[0]
		x := yx[1]

		if c[y-1][x] == "." && dist[y-1][x] < 0 {
			dist[y-1][x] = dist[y][x] + 1

			next := []int{y - 1, x}
			q.push(next)
		}
		if c[y+1][x] == "." && dist[y+1][x] < 0 {
			dist[y+1][x] = dist[y][x] + 1

			next := []int{y + 1, x}
			q.push(next)
		}
		if c[y][x-1] == "." && dist[y][x-1] < 0 {
			dist[y][x-1] = dist[y][x] + 1

			next := []int{y, x - 1}
			q.push(next)
		}
		if c[y][x+1] == "." && dist[y][x+1] < 0 {
			dist[y][x+1] = dist[y][x] + 1

			next := []int{y, x + 1}
			q.push(next)
		}
	}

	fmt.Println(dist[g[0]][g[1]])
}
