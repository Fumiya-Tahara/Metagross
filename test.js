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