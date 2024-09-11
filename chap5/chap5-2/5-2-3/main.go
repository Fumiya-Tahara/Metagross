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

func main() {
	sc.Buffer(buffer, maxVal)

	NK := scanLineInt64()
	N := NK[0]
	K := NK[1]

	A := scanLineInt64()
	A = append([]int64{0}, A...)

	First := make([]int, N+1)
	Second := make([]int, N+1)

	for i := 1; i <= int(N); i++ {
		First[i] = -1
		Second[i] = -1
	}

	count := 0
	var current int64 = 1
	for {
		if First[current] == -1 {
			First[current] = count
		} else if Second[current] == -1 {
			Second[current] = count
		}

		if int64(count) == K {
			fmt.Println(current)
			return
		} else if Second[current] != -1 && (K-int64(First[current]))%int64(Second[current]-First[current]) == 0 {
			fmt.Println(current)
			return
		}

		current = A[current]
		count++
	}

}
