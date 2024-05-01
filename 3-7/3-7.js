// 等差数列
function p1(n) {
  if (n <= 1) {
    return n;
  } else {
    return p1(n - 1) + 2;
  }
}
console.log(p1(4));

//フィボナッチ数列
function p2(n) {
  if (n <= 1) {
    return n;
  } else {
    return p2(n - 1) + p2(n - 2);
  }
}
console.log(p2(4));

//さらに複雑な漸化式
let h = [8, 6, 9, 2, 1];
let a = [0, 2];

for (let n = 2; n < 5; n++) {
  a[n] = Math.min(
    a[n - 1] + Math.abs(h[n - 1] - h[n]),
    a[n - 2] + Math.abs(h[n - 2] - h[n])
  );
}

console.log(a[4]);

//カエルの移動の一般化
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