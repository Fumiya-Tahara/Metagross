function main() {
  const N = 1000;
  let answer = 0.0;

  for (let i = 0; i < N; i++) {
    let x = (2 * i + 1) / (2 * N);
    let value = Math.pow(2.0, x * x);
    answer += value;
  }

  console.log((answer / N).toFixed(14));//小数点以下14桁
}

main();//1.28822624878143
