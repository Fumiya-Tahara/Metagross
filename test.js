const readline = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
});

readline.question('階段の段数を入力してください: ', input => {
  let N = Number(input);
  let dp = new Array(N + 1).fill(0);

  dp[0] = 1; // 0段目
  dp[1] = 1; // 1段目

  for (let n = 2; n <= N; n++) {
    dp[n] = dp[n-1] + dp[n-2];
  }

  console.log(dp[N]);
  readline.close();
});