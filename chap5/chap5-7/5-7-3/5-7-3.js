function Main(input) {
  const lines = input.trim().split("\n");
  const N = parseInt(lines[0], 10);
  const A = lines[1].split(" ").map(BigInt);

  A.sort((a, b) => (a > b ? 1 : -1));

  let Answer = 0n;
  for (let i = 1; i <= N; i++) {
    Answer += A[i - 1] * BigInt(-N + 2 * i - 1);
  }

  console.log(Answer.toString());
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
