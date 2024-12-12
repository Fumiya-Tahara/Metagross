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

func dfs(g *[][]int, visited *[]bool, dp *[]int, pos int) {
	(*visited)[pos] = true
	(*dp)[pos] = 1
	for _, v := range (*g)[pos] {
		if (*visited)[v] {
			continue
		}
		dfs(g, visited, dp, v)
		(*dp)[pos] += (*dp)[v]
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()
	g := make([][]int, n+1)
	for i := 0; i < n; i++ {
		a := scan()
		b := scan()
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	dp := make([]int, n+1)
	visited := make([]bool, n+1)

	dfs(&g, &visited, &dp, 1)

	ans := 0
	for i := 2; i <= n; i++ {
		ans += dp[i] * (n - dp[i])
	}

	fmt.Println(ans)
}
