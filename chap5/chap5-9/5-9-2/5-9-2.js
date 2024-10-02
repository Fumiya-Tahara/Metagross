function Main(input) {
  const lines = input.trim().split("\n");
  const N = parseInt(lines[0], 10);
  const A = lines[1].split(" ").map(BigInt);
  const B = lines[2].split(" ").map(BigInt);

  A.sort((a, b) => (a > b ? 1 : -1));
  B.sort((a, b) => (a > b ? 1 : -1));

  let totalDistance = 0n;
  for (let i = 0; i < N; i++) {
    totalDistance += A[i] > B[i] ? A[i] - B[i] : B[i] - A[i];
  }

  console.log(totalDistance.toString());
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));
