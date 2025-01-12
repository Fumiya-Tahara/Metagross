package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if 1e18/(a/gcd(a, b)) < b {
		fmt.Println("Large")
		return
	}
	fmt.Println(lcm(a, b))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
