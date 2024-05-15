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

//階段の上り方
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

//ナップザック問題
const readline = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
});

let inputIndex = 0;
let inputs = [];

readline.on('line', (input) => {
  inputs.push(input);
  inputIndex++;

  if (inputIndex - 1 < Number(inputs[0].split(' ')[0])) {
    console.log(`品物${inputIndex}の重さと価値をスペース区切りで入力: `);
  } else if (inputIndex - 1 === Number(inputs[0].split(' ')[0])) {
    let [N, W] = inputs[0].split(' ').map(Number);
    let weights = inputs.slice(1).map(item => Number(item.split(' ')[0]));
    let values = inputs.slice(1).map(item => Number(item.split(' ')[1]));
    let dp = new Array(W + 1).fill(0);
		console.log(dp);

    for (let i = 0; i < N; i++) {//品物の数だけループ
      for (let j = W; j >= weights[i]; j--) {//重さの上限から品物の重さだけ引いた数だけループ
				console.log(j, weights[i])
				console.log(dp[j], dp[j - weights[i]], values[i])
        dp[j] = Math.max(dp[j], dp[j - weights[i]] + values[i]);//dp[j]とdp[j - weights[i]] + values[i]の大きい方をdp[j]に代入
				console.log(dp);
      }
    }

    console.log(dp[W]);
    readline.close();
  }
});

console.log('品物の数Nと重さの上限Wをスペース区切りで入力: ');

//再帰関数
const readline = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
});

readline.question('品物の数Nと重さの上限Wをスペース区切りで入力: ', input => {
  let [N, W] = input.split(' ').map(Number);
  let weights = new Array(N);
  let values = new Array(N);

  const readItems = (index = 0) => {
    if (index < N) {
      readline.question(`品物${index+1}の重さと価値をスペース区切りで入力: `, item => {
        let [w, v] = item.split(' ').map(Number);
        weights[index] = w;
        values[index] = v;
        readItems(index + 1);
      });
    } else {
      let dp = new Array(W + 1).fill(0);

      for (let i = 0; i < N; i++) {
        for (let j = W; j >= weights[i]; j--) {
          dp[j] = Math.max(dp[j], dp[j - weights[i]] + values[i]);
        }
      }

      console.log(dp[W]);
      readline.close();
    }
  };

  readItems();
});

//節末問題
//3.7.4
const readline = require('readline');

let rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.question('カードの数Nと目標の合計Sをスペース区切りで入力: ', input => {
  let [N, S] = input.split(' ').map(Number);
  let cards = new Array(N);

  const readCards = (index = 0) => {
    if (index < N) {
      rl.question(`カード${index+1}の数値を入力してください: `, card => {
        cards[index] = Number(card);
        readCards(index + 1);
      });
    } else {
      let dp = new Array(S + 1).fill(false);
      dp[0] = true;

      for (let i = 0; i < N; i++) {
        for (let j = S; j >= cards[i]; j--) {
          dp[j] = dp[j] || dp[j - cards[i]];
        }
      }

      console.log(dp[S] ? 'Yes' : 'No');
      rl.close();
    }
  };

  readCards();
});

//3.7.6
const readline = require('readline');

let rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.question('夏休みの日数Nを入力してください: ', input => {
  let N = Number(input);
  let powerUps = new Array(N);

  const readDays = (index = 0) => {//再帰関数
    if (index < N) {
      rl.question(`日${index+1}に勉強すると実力がいくつ上がるかを入力してください: `, powerUp => {
        powerUps[index] = Number(powerUp);
        readDays(index + 1);//再帰関数
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