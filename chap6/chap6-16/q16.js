function Main(input) {
  const [N, X] = input.trim().split(" ").map(Number);
  let count = 0;

  for (let a = 1; a <= N - 2; a++) {
    for (let b = a + 1; b <= N - 1; b++) {
      for (let c = b + 1; c <= N; c++) {
        if (a + b + c === X) {
          count++;
        }
      }
    }
  }

  console.log(count);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));