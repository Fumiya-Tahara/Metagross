function Main(input) {
  const lines = input.trim().split("\n");
  const N = parseInt(lines[0], 10);

  const adjList = Array.from({ length: N + 1 }, () => []);
  for (let i = 1; i < N; i++) {
    const [a, b] = lines[i].split(" ").map(Number);
    adjList[a].push(b);
    adjList[b].push(a);
  }

  function dfs(node) {
    visited[node] = true;
    dp[node] = 1;
    for (const neighbor of adjList[node]) {
      if (!visited[neighbor]) {
        dfs(neighbor, node);
        dp[node] += dp[neighbor];
      }
    }
  }

  const visited = Array(N + 1).fill(false);
  const dp = Array(N + 1).fill(0);

  dfs(1, -1);

  let answer = 0;
  for (let i = 1; i <= N; i++) {
    answer += dp[i] * (N - dp[i]);
  }
  console.log(answer);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
