package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	return n
}

func dfs(current int, g *[][]int, visited *[]bool, dp *[]int) {
	(*visited)[current] = true
	(*dp)[current] = 1

	for _, v := range (*g)[current] {
		if !(*visited)[v] {
			dfs(v, g, visited, dp)
			(*dp)[current] += (*dp)[v]
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()

	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		a := scan()
		b := scan()
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	visited := make([]bool, n+1)
	dp := make([]int, n+1)

	dfs(1, &g, &visited, &dp)

	var answer int
	for i := 1; i < n+1; i++ {
		answer += dp[i] * (n - dp[i])
	}

	fmt.Println(answer)
}
