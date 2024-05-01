const readline = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
});

readline.question('足場の高さを5つ、カンマ区切りで入力してください: ', input => {
  let h = input.split(',').map(Number);
  let dp = new Array(h.length).fill(Infinity);

  dp[0] = 0;

  for (let i = 1; i < h.length; i++) {
    if (i == 1) {
      dp[i] = Math.abs(h[i] - h[i-1]);
    } else {
      dp[i] = Math.min(dp[i-1] + Math.abs(h[i] - h[i-1]), dp[i-2] + Math.abs(h[i] - h[i-2]));
    }
  }

  console.log(dp[h.length - 1]);
  readline.close();
});