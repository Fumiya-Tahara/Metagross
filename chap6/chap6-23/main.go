package main

import "fmt"

func main() {
	var l, r int64
	fmt.Scan(&l, &r)

	isPrime := make([]bool, r-l+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	if l == 1 {
		isPrime[0] = false
	}

	for i := int64(2); i*i <= r; i++ {
		minVal := ((l + i - 1) / i) * i
		for j := minVal; j <= r; j += i {
			if j == i {
				continue
			}
			isPrime[j-l] = false
		}
	}

	var answer int
	for i := int64(0); i <= r-l; i++ {
		if isPrime[i] {
			answer += 1
		}
	}

	fmt.Println(answer)
}
