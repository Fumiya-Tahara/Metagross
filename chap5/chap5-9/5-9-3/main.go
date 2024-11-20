package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

type Movie struct {
	l, r int
}

type ByEndTime []Movie

func (a ByEndTime) Len() int      { return len(a) }
func (a ByEndTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByEndTime) Less(i, j int) bool {
	if a[i].r < a[j].r {
		return true
	}
	if a[i].r > a[j].r {
		return false
	}
	return a[i].l < a[j].l
}

func main() {
	sc.Buffer(buffer, maxVal)

	N := scanInt()

	LR := make([]Movie, N)
	for i := 0; i < N; i++ {
		lr := scanLineInt()
		LR[i] = Movie{l: lr[0], r: lr[1]}
	}

	sort.Sort(ByEndTime(LR))

	CurrentTime := 0
	Answer := 0
	for i := 0; i < N; i++ {
		if CurrentTime <= LR[i].l {
			CurrentTime = LR[i].r
			Answer++
		}
	}

	fmt.Println(Answer)
}
