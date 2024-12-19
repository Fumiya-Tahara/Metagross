function Main(input) {
  const lines = input.trim().split("\n");
  const N = parseInt(lines[0], 10);
  const T = lines[1].split(" ").map(Number);

  let sumT = 0;
  for (let i = 0; i < N; i++) {
    sumT += T[i];
  }

  const dp = Array.from({ length: N + 1 }, () => Array(sumT + 1).fill(false));
  dp[0][0] = true;

  for (let i = 1; i <= N; i++) {
    for (let j = 0; j <= sumT; j++) {
      if (j < T[i - 1]) {
        dp[i][j] = dp[i - 1][j];
      } else {
        dp[i][j] = dp[i - 1][j] || dp[i - 1][j - T[i - 1]];
      }
    }
  }

  let answer = Infinity;
  for (let i = 0; i <= sumT; i++) {
    if (dp[N][i]) {
      const cookingTime = Math.max(i, sumT - i);
      answer = Math.min(answer, cookingTime);
    }
  }
  console.log(answer);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
