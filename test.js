const readline = require('readline');

let rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.question('夏休みの日数Nを入力してください: ', input => {
  let N = Number(input);
  let powerUps = new Array(N);

  const readDays = (index = 0) => {
    if (index < N) {
      rl.question(`日${index+1}に勉強すると実力がいくつ上がるかを入力してください: `, powerUp => {
        powerUps[index] = Number(powerUp);
        readDays(index + 1);
      });
    } else {
      let dp = new Array(N + 1).fill(0);

      for (let i = 0; i < N; i++) {
        dp[i + 1] = Math.max(dp[i + 1], dp[i]);
        if (i + 2 <= N) {
          dp[i + 2] = Math.max(dp[i + 2], dp[i] + powerUps[i]);
        }
      }

      console.log(dp[N]); // 夏休みの間に実力がいくつ上がるかの最大値を出力
      rl.close();
    }
  };

  readDays();
});