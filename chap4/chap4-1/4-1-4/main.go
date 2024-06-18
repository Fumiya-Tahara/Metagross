package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	maxN   = 2000
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
	sc.Buffer(buffer, maxN*12)

	abhm := scanLineInt()
	A := abhm[0]
	B := abhm[1]
	H := abhm[2]
	M := abhm[3]

	hourHandF := 12.0 * 60.0
	minuteHandF := 60.0

	hourxy := make([]float64, 2)
	minutexy := make([]float64, 2)

	hourAngle := float64(H)*60.0 + float64(M)
	minuteAngle := float64(M)

	hourxy[0] = float64(A) * math.Sin(2.0*math.Pi*hourAngle/hourHandF)
	hourxy[1] = float64(A) * math.Cos(2.0*math.Pi*hourAngle/hourHandF)

	minutexy[0] = float64(B) * math.Sin(2.0*math.Pi*minuteAngle/minuteHandF)
	minutexy[1] = float64(B) * math.Cos(2.0*math.Pi*minuteAngle/minuteHandF)

	dist := math.Sqrt(math.Pow(hourxy[0]-minutexy[0], 2) + math.Pow(hourxy[1]-minutexy[1], 2))
	fmt.Println(dist)
}
