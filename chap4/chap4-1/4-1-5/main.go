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

func cross(a, b, c Point) int {
	return (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
}

type Point struct {
	X, Y int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	sc.Buffer(buffer, maxN*12)

	xy1 := scanLineInt()
	xy2 := scanLineInt()
	xy3 := scanLineInt()
	xy4 := scanLineInt()

	A := Point{xy1[0], xy1[1]}
	B := Point{xy2[0], xy2[1]}
	C := Point{xy3[0], xy3[1]}
	D := Point{xy4[0], xy4[1]}

	// ベクトルAB・ベクトルACの外積、ベクトルAB・ベクトルADの外積、ベクトルCD・ベクトルCAの外積、ベクトルCD・ベクトルCBの外積
	ans1 := cross(A, B, C)
	ans2 := cross(A, B, D)
	ans3 := cross(C, D, A)
	ans4 := cross(C, D, B)

	// ABCDが一直線上にある場合
	if ans1 == 0 && ans2 == 0 && ans3 == 0 && ans4 == 0 {
		var isCross bool
		if A.X == B.X && A.X == C.X && A.X == D.X {
			if A.Y > B.Y {
				A, B = B, A
			}
			if C.Y > D.Y {
				C, D = D, C
			}
			if max(A.Y, C.Y) <= min(B.Y, D.Y) {
				isCross = true
			}
		} else {
			if A.X > B.X {
				A, B = B, A
			}
			if C.X > D.X {
				C, D = D, C
			}
			if max(A.X, C.X) <= min(B.X, D.X) {
				isCross = true
			}
		}
		if isCross {
			fmt.Println("Yes")
			return
		}
		fmt.Println("No")
		return
	}

	// ABCDが一直線上にない場合
	isAB := false
	isCD := false
	if ans1 >= 0 && ans2 <= 0 {
		isAB = true
	}
	if ans1 <= 0 && ans2 >= 0 {
		isAB = true
	}
	if ans3 >= 0 && ans4 <= 0 {
		isCD = true
	}
	if ans3 <= 0 && ans4 >= 0 {
		isCD = true
	}

	if isAB && isCD {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")

}
