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

// 行列式を使って交点を求める
func findIntersection(abc1, abc2 []float64) (float64, float64) {
	det := abc1[0]*abc2[1] - abc2[0]*abc1[1]
	if det == 0 {
		log.Println("平行")
		return 0, 0
	}

	x := (abc1[2]*abc2[1] - abc2[2]*abc1[1]) / det
	y := (abc1[0]*abc2[2] - abc2[0]*abc1[2]) / det

	if x == -0 {
		x = 0
	}

	if y == -0 {
		y = 0
	}

	return x, y
}

func main() {
	sc.Buffer(buffer, maxVal)

	N := scanInt()
	abc := make([][]float64, N)

	for i := 0; i < N; i++ {
		row := scanLineFloat64()
		abc[i] = row
	}

	var answer float64
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			x, y := findIntersection(abc[i], abc[j])

			valid := true
			for k := 0; k < N; k++ {
				check := abc[k][0]*x + abc[k][1]*y
				if check > abc[k][2] {
					valid = false
					break
				}
			}
			if valid && x+y > answer {
				answer = x + y
			}
		}
	}

	fmt.Println(answer)
}
